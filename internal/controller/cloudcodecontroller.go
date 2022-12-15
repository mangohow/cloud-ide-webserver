package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/model/reqtype"
	"github.com/mangohow/cloud-ide-webserver/internal/service"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/serialize"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type CloudCodeController struct {
	logger       *logrus.Logger
	spaceService *service.CloudCodeService
}

func NewCloudCodeController() *CloudCodeController {
	return &CloudCodeController{
		logger:       logger.Logger(),
		spaceService: service.NewCloudCodeService(),
	}
}

// CreateSpace 创建一个云空间  method: POST path: /api/space
// Request Param: reqtype.SpaceCreateOption
func (c *CloudCodeController) CreateSpace(ctx *gin.Context) *serialize.Response {
	// 1、用户参数获取和验证
	req := c.creationCheck(ctx)
	if req == nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	// 2、获取用户id，在token验证时已经解析出并放入ctx中了
	idi, _ := ctx.Get("id")
	id := idi.(uint32)

	// 3、调用service处理然后响应结果
	space, err := c.spaceService.CreateWorkspace(req, id)
	switch err {
	case service.ErrNameDuplicate:
		return serialize.NewResponseOKND(code.SpaceCreateNameDuplicate)
	case service.ErrReachMaxSpaceCount:
		return serialize.NewResponseOKND(code.SpaceCreateReachMaxCount)
	case service.ErrCreate:
		return serialize.NewResponseOKND(code.SpaceCreateFailed)
	case service.ErrReqParamInvalid:
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	if err != nil {
		return serialize.NewResponseOKND(code.SpaceCreateFailed)
	}

	return serialize.NewResponseOk(code.SpaceCreateSuccess, space)
}

// creationCheck 用户参数验证
func (c *CloudCodeController) creationCheck(ctx *gin.Context) *reqtype.SpaceCreateOption {
	// 获取用户请求参数
	var req reqtype.SpaceCreateOption
	// 绑定数据
	err := ctx.ShouldBind(&req)
	if err != nil {
		return nil
	}

	c.logger.Debug(req)

	// 参数验证
	get1, exist1 := ctx.Get("id")
	_, exist2 := ctx.Get("username")
	if !exist1 || !exist2 {
		return nil
	}
	id, ok := get1.(uint32)
	if !ok || id != req.UserId {
		return nil
	}

	return &req
}

// CreateSpaceAndStart 创建一个新的云空间并启动 method: POST path: /api/space_cas
// Request Param: reqtype.SpaceCreateOption
func (c *CloudCodeController) CreateSpaceAndStart(ctx *gin.Context) *serialize.Response {
	req := c.creationCheck(ctx)
	if req == nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	idi, _ := ctx.Get("id")
	id := idi.(uint32)
	uidi, _ := ctx.Get("uid")
	uid := uidi.(string)

	space, err := c.spaceService.CreateAndStartWorkspace(req, id, uid)
	switch err {
	case service.ErrNameDuplicate:
		return serialize.NewResponseOKND(code.SpaceCreateNameDuplicate)
	case service.ErrReachMaxSpaceCount:
		return serialize.NewResponseOKND(code.SpaceCreateReachMaxCount)
	case service.ErrCreate:
		return serialize.NewResponseOKND(code.SpaceCreateFailed)
	case service.ErrSpaceStart:
		return serialize.NewResponseOKND(code.SpaceStartFailed)
	case service.ErrReqParamInvalid:
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	if err != nil {
		return serialize.NewResponseOKND(code.SpaceCreateFailed)
	}

	return serialize.NewResponseOk(code.SpaceStartSuccess, space)
}

func (c *CloudCodeController) generatePodName(username string, id uint32) string {
	return username + strconv.Itoa(int(id)) + "-" + strconv.Itoa(int(time.Now().Unix()))
}

// StartSpace 启动一个已存在的云空间
func (c *CloudCodeController) StartSpace(ctx *gin.Context) *serialize.Response {
	panic("TODO")
}

// StopSpace 停止正在运行的云空间 method: PUT path: /api/space_stop
// Request Param: sid
func (c *CloudCodeController) StopSpace(ctx *gin.Context) *serialize.Response {
	var req struct{
		Sid string `json:"sid"`
	}
	err := ctx.ShouldBind(&req)
	if err != nil {
		c.logger.Warningf("bind error:%v", err)
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	uidi, ok := ctx.Get("uid")
	if !ok {
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	uid := uidi.(string)
	err = c.spaceService.StopWorkspace(req.Sid, uid)
	if err != nil {
		if err == service.ErrWorkSpaceIsNotRunning {
			return serialize.NewResponseOKND(code.SpaceStopIsNotRunning)
		}

		return serialize.NewResponseOKND(code.SpaceStopFailed)
	}

	return serialize.NewResponseOKND(code.SpaceStopSuccess)
}

// DeleteSpace 删除已存在的云空间  method: DELETE path: /api/delete
// Request Param: id
func (c *CloudCodeController) DeleteSpace(ctx *gin.Context) *serialize.Response {
	id, err := utils.QueryUint32(ctx, "id")
	if err != nil {
		c.logger.Warningf("get param sid failed:%v", err)
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	c.logger.Debug("id:", id)

	err = c.spaceService.DeleteWorkspace(id)
	if err != nil {
		if err == service.ErrWorkSpaceIsRunning {
			return serialize.NewResponseOKND(code.SpaceDeleteIsRunning)
		}

		return serialize.NewResponseOKND(code.SpaceDeleteFailed)
	}

	return serialize.NewResponseOKND(code.SpaceDeleteSuccess)
}

// ListSpace 获取所有创建的云空间 method: GET path: /api/spaces
// Request param: id uid
func (c *CloudCodeController) ListSpace(ctx *gin.Context) *serialize.Response {
	v1, e1 := ctx.Get("id")
	v2, e2 := ctx.Get("uid")
	if !e1 || !e2 {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	id := v1.(uint32)
	uid := v2.(string)

	spaces, err := c.spaceService.ListWorkspace(id, uid)
	if err != nil {
		return serialize.NewResponseOKND(code.QueryFailed)
	}

	return serialize.NewResponseOk(code.QuerySuccess, spaces)
}
