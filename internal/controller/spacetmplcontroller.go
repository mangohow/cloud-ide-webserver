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
	tmpls, err := s.service.GetAllUsingTmpl()
	if err != nil {
		s.logger.Warnf("get tmpls err:%v", err)
		return serialize.NewResponseOk(code.QueryFailed, nil)
	}

	return serialize.NewResponseOk(code.QuerySuccess, tmpls)
}
