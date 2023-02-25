package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/service"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/serialize"
	"github.com/sirupsen/logrus"
)

type SpaceTmplController struct {
	logger  *logrus.Logger
	service *service.SpaceTmplService
}

func NewSpaceTmplController() *SpaceTmplController {
	return &SpaceTmplController{
		logger:  logger.Logger(),
		service: service.NewSpaceTmplService(),
	}
}

// SpaceTmpls 获取所有模板 method: GET path:/api/tmpls
func (s *SpaceTmplController) SpaceTmpls(ctx *gin.Context) *serialize.Response {
	tmpls, kinds, err := s.service.GetAllUsingTmpl()
	if err != nil {
		s.logger.Warnf("get tmpls err:%v", err)
		return serialize.NewResponseOKND(code.QueryFailed)
	}

	return serialize.NewResponseOk(code.QuerySuccess, gin.H{
		"tmpls": tmpls,
		"kinds": kinds,
	})
}

// SpaceSpecs 获取空间规格 method: GET path:/api/specs
func (s *SpaceTmplController) SpaceSpecs(ctx *gin.Context) *serialize.Response {
	specs, err := s.service.GetAllSpec()
	if err != nil {
		s.logger.Warnf("get specs error:%v", err)
		return serialize.NewResponseOKND(code.QueryFailed)
	}

	return serialize.NewResponseOk(code.QuerySuccess, specs)
}