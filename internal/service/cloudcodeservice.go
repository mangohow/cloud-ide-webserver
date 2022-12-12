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
	factory := caches.NewCacheFactory()
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
	ErrCreate             = errors.New("create failed")
	ErrSpaceStart         = errors.New("space start failed")
)

// CreateWorkspace 创建云工作空间
func (c *CloudCodeService) CreateWorkspace(req *reqtype.SpaceCreateOption, userId uint32) (*model.Space, error) {
	// 1、验证创建的工作空间是否达到最大数量
	count, err := c.dao.FindCountByUserId(userId)
	if err != nil {
		c.logger.Warnf("get space count error:%v", err)
		return nil, ErrCreate
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
	tmpl := c.tmplCache.Get(req.TmplId)
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
		Sid: generateSID(),
	}

	//6、 添加到数据库
	spaceId, err := c.dao.Insert(space)
	if err != nil {
		c.logger.Errorf("add space error:%v", err)
		return nil, ErrCreate
	}
	space.Id = spaceId

	return space, nil
}

// CreateAndStartWorkspace 创建并且启动云工作空间
func (c *CloudCodeService) CreateAndStartWorkspace(req *reqtype.SpaceCreateOption, userId uint32, uid string) (*model.Space, error) {
	// TODO 检查是否有工作空间正在运行, 需要停止
	// TODO 修改controller中运行时的Pod名称，不能根据用户指定的工作空间名称来生成Pod名称，因为Pod名称中不能含有'_'等字符

	// 1、创建工作空间
	space, err := c.CreateWorkspace(req, userId)
	if err != nil {
		return nil, err
	}

	// 2、获取模板
	tmpl := c.tmplCache.Get(req.TmplId)
	if tmpl == nil {
		c.logger.Warnf("get tmpl cache error:%v", err)
		return nil, ErrCreate
	}

	// 3、生成Pod名称
	podName := c.generatePodName(space.Sid, uid)

	pod := pb.PodInfo{
		Name:      podName,
		Namespace: CloudCodeNamespace,
		Image:     tmpl.Image,
		Port:      DefaultPodPort,
		ResourceLimit: &pb.ResourceLimit{
			Cpu:     space.Spec.CpuSpec,
			Memory:  space.Spec.MemSpec,
			Storage: space.Spec.StorageSpec,
		},
	}

	// 5、请求k8s controller创建云空间
	spaceInfo, err := c.rpc.CreateSpaceAndWaitForRunning(context.Background(), &pod)
	if err != nil {
		c.logger.Warnf("rpc create space and wait error:%v", err)
		return nil, ErrSpaceStart
	}

	// 访问路径为  http://domain/ws/uid/...   ws: workspace
	// 7、将相关信息保存到redis
	host := spaceInfo.Ip + ":" + strconv.Itoa(int(spaceInfo.Port))
	err = rdis.AddRunningSpace(uid, &model.RunningSpace{
		Sid:  space.Sid,
		Uid:  space.Name,
		Host: host,
	})
	if err != nil {
		c.logger.Errorf("add pod info to redis error, err:%v", err)
		return nil, ErrSpaceStart
	}

	space.RunningStatus = model.RunningStatusRunning

	return space, nil
}

var (
	ErrWorkSpaceIsRunning = errors.New("workspace is running")
	ErrWorkSpaceIsNotRunning = errors.New("workspace is not running")
)

// DeleteWorkspace 删除云工作空间
func (c *CloudCodeService) DeleteWorkspace(id uint32) error {
	// 1、检查该工作空间是否正在运行，如果正在运行就返回错误
	sid, err := c.dao.FindSidById(id)
	if err != nil {
		c.logger.Warnf("find sid error:%v", err)
		return err
	}
	// 从redis中查询
	isRunning, err := rdis.CheckIsRunning(sid)
	if err != nil {
		c.logger.Warnf("check is running error:%v", err)
		return err
	}
	if isRunning {
		return ErrWorkSpaceIsRunning
	}

	// 2、从mysql中删除记录
	return c.dao.DeleteSpaceById(id)
}

// StopWorkspace 停止云工作空间
func (c *CloudCodeService) StopWorkspace(id uint32, sid, uid string) error {
	// 1、查询云工作空间是否正在运行
	isRunning, err := rdis.CheckIsRunning(sid)
	if err != nil {
		c.logger.Warnf("check is running error:%v", err)
		return err
	}
	if !isRunning {
		return ErrWorkSpaceIsNotRunning
	}

	// 2、停止workspace
	name := c.generatePodName(sid, uid)
	_, err = c.rpc.DeleteSpace(context.Background(), &pb.QueryOption{
		Name:      name,
		Namespace: CloudCodeNamespace,
	})
	if err != nil {
		c.logger.Warnf("rpc delete space error:%v", err)
		return err
	}

	return nil
}

// StartWorkspace 启动云工作空间
func (c *CloudCodeService) StartWorkspace() {

}

// ListWorkspace 列出云工作空间
func (c *CloudCodeService) ListWorkspace(userId uint32, uid string) ([]model.Space, error) {
	spaces, err := c.dao.FindAllSpaceByUserId(userId)
	if err != nil {
		c.logger.Warnf("find spaces error:%v", err)
		return nil, err
	}

	runningSpace, err := rdis.GetRunningSpace(uid)
	if err != nil {
		c.logger.Warnf("get running space error:%v", err)
		return spaces, nil
	}
	if runningSpace == nil {
		return spaces, nil
	}

	for i, item := range spaces {
		if item.Name == runningSpace.Uid {
			spaces[i].RunningStatus = model.RunningStatusRunning
		}
	}

	return spaces, nil
}

// generatePodName 生成Pod的名称
// podName: "ws"-uid-sid
func (c *CloudCodeService) generatePodName(sid, uid string) string {
	return strings.Join([]string{"ws", uid, sid}, "-")
}

// generateSID 生成Space id
func generateSID() string {
	return bson.NewObjectId().Hex()
}
