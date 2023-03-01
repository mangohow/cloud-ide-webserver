package service

import (
	"context"
	"errors"
	"github.com/mangohow/cloud-ide-webserver/internal/caches"
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/dao/rdis"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/internal/model/reqtype"
	"github.com/mangohow/cloud-ide-webserver/internal/rpc"
	"github.com/mangohow/cloud-ide-webserver/pb"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

const (
	CloudCodeNamespace = "cloud-ide"
	DefaultPodPort     = 9999
	MaxSpaceCount      = 20
)

type CloudCodeService struct {
	logger    *logrus.Logger
	rpc       pb.CloudIdeServiceClient
	dao       *dao.SpaceDao
	tmplCache *caches.TmplCache
	specCache *caches.SpecCache
}

func NewCloudCodeService() *CloudCodeService {
	conn := rpc.GrpcClient("space-code")
	factory := caches.CacheFactory()
	d := dao.NewSpaceTemplateDao()
	return &CloudCodeService{
		logger:    logger.Logger(),
		rpc:       pb.NewCloudIdeServiceClient(conn),
		dao:       dao.NewSpaceDao(),
		tmplCache: factory.TmplCache(d),
		specCache: factory.SpecCache(d),
	}
}

var (
	ErrReqParamInvalid    = errors.New("request param invalid")
	ErrNameDuplicate      = errors.New("name duplicate")
	ErrReachMaxSpaceCount = errors.New("reach max space count")
	ErrSpaceCreate        = errors.New("space create failed")
	ErrSpaceStart         = errors.New("space start failed")
	ErrSpaceAlreadyExist  = errors.New("space already exist")
	ErrSpaceNotFound      = errors.New("space not found")
	ErrResourceExhausted  = errors.New("no adequate resource are available")
)

// CreateWorkspace 创建云工作空间, 只涉及数据库操作
func (c *CloudCodeService) CreateWorkspace(req *reqtype.SpaceCreateOption, userId uint32) (*model.Space, error) {
	// 1、验证创建的工作空间是否达到最大数量
	count, err := c.dao.FindCountByUserId(userId)
	if err != nil {
		c.logger.Warnf("get space count error:%v", err)
		return nil, ErrSpaceCreate
	}
	if count >= MaxSpaceCount {
		return nil, ErrReachMaxSpaceCount
	}

	// 2、验证名称是否重复
	if err := c.dao.FindByUserIdAndName(userId, req.Name); err == nil {
		c.logger.Warnf("find space error:%v", err)
		return nil, ErrNameDuplicate
	}

	// 3、从缓存中获取要创建的云空间的模板
	tmpl := c.tmplCache.GetTmpl(req.TmplId)
	if tmpl == nil {
		c.logger.Warnf("get tmpl cache error:%v", err)
		return nil, ErrReqParamInvalid
	}

	// 4、从缓存中获取要创建的云空间的规格
	spec := c.specCache.Get(req.SpaceSpecId)
	if spec == nil {
		return nil, ErrReqParamInvalid
	}

	// 5、构造云工作空间结构
	now := time.Now()
	space := &model.Space{
		UserId:     userId,
		TmplId:     tmpl.Id,
		SpecId:     spec.Id,
		Spec:       *spec,
		Name:       req.Name,
		Status:     model.SpaceStatusUncreated,
		CreateTime: now,
		DeleteTime: now,
		StopTime:   now,
		TotalTime:  0,
		Sid:        generateSID(),
	}

	//6、 添加到数据库
	spaceId, err := c.dao.Insert(space)
	if err != nil {
		c.logger.Errorf("add space error:%v", err)
		return nil, ErrSpaceCreate
	}
	space.Id = spaceId

	return space, nil
}

var ErrOtherSpaceIsRunning = errors.New("there is other space running")

// CreateAndStartWorkspace 创建并且启动云工作空间
func (c *CloudCodeService) CreateAndStartWorkspace(req *reqtype.SpaceCreateOption, userId uint32, uid string) (*model.Space, error) {
	// 1、检查是否有其它工作空间正在运行, 同时只能有一个工作空间启动
	isRunning, err := rdis.CheckHasRunningSpace(uid)
	if err != nil {
		return nil, ErrSpaceCreate
	}
	if isRunning {
		return nil, ErrOtherSpaceIsRunning
	}

	// 2、创建工作空间
	space, err := c.CreateWorkspace(req, userId)
	if err != nil {
		return nil, err
	}

	// 3、启动工作空间
	return c.startWorkspace(space, uid, c.rpc.CreateSpace)
}

type StartFunc func(ctx context.Context, in *pb.WorkspaceInfo, opts ...grpc.CallOption) (*pb.WorkspaceRunningInfo, error)

// startWorkspace 启动工作空间
func (c *CloudCodeService) startWorkspace(space *model.Space, uid string, startFunc StartFunc) (*model.Space, error) {
	// 1、获取空间模板
	tmpl := c.tmplCache.GetTmpl(space.TmplId)
	if tmpl == nil {
		c.logger.Warnf("get tmpl cache error")
		return nil, ErrSpaceStart
	}

	// 2、生成Pod信息
	podName := c.generatePodName(space.Sid, uid)
	pod := pb.WorkspaceInfo{
		Name:      podName,
		Namespace: CloudCodeNamespace,
		Image:     tmpl.Image,
		Port:      DefaultPodPort,
		VolumeMountPath: "/user_data/",
		ResourceLimit: &pb.ResourceLimit{
			Cpu:     space.Spec.CpuSpec,
			Memory:  space.Spec.MemSpec,
			Storage: space.Spec.StorageSpec,
		},
	}

	var retErr error
loop:
	for i := 0; i < 1; i++ {
		// 3、请求k8s controller创建云空间
		// 设置一分钟的超时时间
		timeout, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
		defer cancelFunc()
		spaceInfo, err := startFunc(timeout, &pod)
		if err != nil {
			s, ok := status.FromError(err)
			if !ok {
				return nil, err
			}

			switch s.Code() {
			// 创建工作空间时,工作空间已存在,修改数据库中的status
			case codes.AlreadyExists:
				retErr = ErrSpaceAlreadyExist
				break loop

			// 启动工作空间时,工作空间不存在
			case codes.NotFound:
				return nil, ErrSpaceNotFound

			// 资源耗尽,无法启动
			case codes.ResourceExhausted:
				return nil, ErrResourceExhausted
			case codes.Unknown:
				c.logger.Errorf("rpc start space error:%v", err)
				return nil, ErrSpaceStart
			}
		}

		// 访问路径为  http://domain/ws/uid/...   ws: workspace
		// 4、将相关信息保存到redis
		host := spaceInfo.Ip + ":" + strconv.Itoa(int(spaceInfo.Port))
		err = rdis.AddRunningSpace(uid, &model.RunningSpace{
			Sid:  space.Sid,
			Host: host,
		})
		if err != nil {
			c.logger.Errorf("add pod info to redis error, err:%v", err)
			return nil, ErrSpaceStart
		}

		space.RunningStatus = model.RunningStatusRunning
	}

	// 5、修改数据库中的状态信息
	if space.Status == model.SpaceStatusUncreated {
		// 更新数据库
		err := c.dao.UpdateStatusById(space.Id, model.SpaceStatusAvailable)
		if err != nil {
			c.logger.Warnf("update space status error:%v", err)
		}
	}

	if retErr != nil {
		return nil, retErr
	}

	return space, nil
}

var (
	ErrWorkSpaceIsRunning    = errors.New("workspace is running")
	ErrWorkSpaceIsNotRunning = errors.New("workspace is not running")
)

// DeleteWorkspace 删除云工作空间
func (c *CloudCodeService) DeleteWorkspace(id uint32, uid string) error {
	// 1、检查该工作空间是否正在运行，如果正在运行就返回错误
	space, err := c.dao.FindSidAndStatusById(id)
	if err != nil {
		c.logger.Warnf("find sid error:%v", err)
		return err
	}
	// 从redis中查询
	isRunning, err := rdis.CheckIsRunning(space.Sid)
	if err != nil {
		c.logger.Warnf("check is running error:%v", err)
		return err
	}
	if isRunning {
		return ErrWorkSpaceIsRunning
	}

	// 2、通知controller删除该workspace联的资源,如果该工作空间没有启动过,就不需要
	if space.Status != model.SpaceStatusUncreated {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
		defer cancelFunc()
		name := c.generatePodName(space.Sid, uid)
		_, err = c.rpc.DeleteSpace(ctx, &pb.QueryOption{Name: name, Namespace: CloudCodeNamespace})
		if err != nil {
			c.logger.Warnf("delete workspace err:%v", err)
			return err
		}
	}

	// 3、从mysql中删除记录
	return c.dao.DeleteSpaceById(id)
}

// StopWorkspace 停止云工作空间
func (c *CloudCodeService) StopWorkspace(sid, uid string) error {
	// 1、查询云工作空间是否正在运行并删除数据
	isRunning, err := rdis.CheckRunningSpaceAndDelete(uid)
	if err != nil {
		c.logger.Warnf("check is running error:%v", err)
		return err
	}
	if !isRunning {
		return ErrWorkSpaceIsNotRunning
	}

	// 2、停止workspace
	name := c.generatePodName(sid, uid)
	_, err = c.rpc.StopSpace(context.Background(), &pb.QueryOption{
		Name:      name,
		Namespace: CloudCodeNamespace,
	})
	if err != nil {
		c.logger.Warnf("rpc delete space error:%v", err)
		return err
	}

	return nil
}

var ErrWorkSpaceNotExist = errors.New("workspace is not exist")

// StartWorkspace 启动云工作空间
func (c *CloudCodeService) StartWorkspace(id, userId uint32, uid string) (*model.Space, error) {
	// 1、检查是否有其它工作空间正在运行, 同时只能有一个工作空间启动
	isRunning, err := rdis.CheckHasRunningSpace(uid)
	if err != nil {
		return nil, ErrSpaceStart
	}
	if isRunning {
		return nil, ErrOtherSpaceIsRunning
	}

	// 2.查询该工作空间是否存在
	space, err := c.dao.FindByIdAndUserId(id, userId)
	if err != nil {
		c.logger.Warnf("find space error:%v", err)
		return nil, ErrWorkSpaceNotExist
	}
	space.Id = id
	space.UserId = userId

	// 3.该工作空间是否是第一次启动
	startFunc := c.rpc.StartSpace
	switch space.Status {
	case model.SpaceStatusDeleted:
		return nil, ErrWorkSpaceNotExist
	case model.SpaceStatusUncreated:
		startFunc = c.rpc.CreateSpace
		spec := c.specCache.Get(space.SpecId)
		if spec == nil {
			return nil, ErrSpaceStart
		}
		space.Spec = *spec
	}

	// 4.启动工作空间
	ret, err := c.startWorkspace(&space, uid, startFunc)
	if err != nil {
		c.logger.Warnf("start workspace error:%v", err)
		return nil, err
	}

	return ret, nil
}

// ListWorkspace 列出云工作空间
func (c *CloudCodeService) ListWorkspace(userId uint32, uid string) ([]model.Space, error) {
	spaces, err := c.dao.FindAllSpaceByUserId(userId)
	if err != nil {
		c.logger.Warnf("find spaces error:%v", err)
		return nil, err
	}

	// 填充environment字段和spec字段
	for i := 0; i < len(spaces); i++ {
		t := c.tmplCache.GetTmpl(spaces[i].TmplId)
		spaces[i].Environment = t.Desc
		spaces[i].Avatar = t.Avatar
		spaces[i].Spec = *c.specCache.Get(spaces[i].SpecId)
		spaces[i].Spec.Id = 0
	}

	runningSpace, err := rdis.GetRunningSpace(uid)
	if err != nil {
		c.logger.Warnf("get running space error:%v, uid:%s", err, uid)
		return spaces, nil
	}

	c.logger.Debugf("ListWorkspace, uid:%s, running space:%v", uid, runningSpace)
	if runningSpace == nil {
		return spaces, nil
	}

	for i, item := range spaces {
		if item.Sid == runningSpace.Sid {
			spaces[i].RunningStatus = model.RunningStatusRunning
			break
		}
	}

	return spaces, nil
}

// generatePodName 生成Pod的名称
// podName: "ws"-uid-sid
func (c *CloudCodeService) generatePodName(sid, uid string) string {
	return strings.Join([]string{"ws", uid, sid}, "-")
}

func (c *CloudCodeService) ModifyName(name string, id, userId uint32) error {
	// 1、验证名称是否重复
	if err := c.dao.FindByUserIdAndName(userId, name); err == nil {
		c.logger.Warnf("find space error:%v", err)
		return ErrNameDuplicate
	}

	// 2.修改名称
	err := c.dao.UpdateNameById(name, id)
	if err != nil {
		c.logger.Warnf("update space name error:%v", err)
		return err
	}

	return nil
}

// generateSID 生成Space id
func generateSID() string {
	return bson.NewObjectId().Hex()
}
