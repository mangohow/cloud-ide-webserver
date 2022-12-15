package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/conf"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
)

func NewGinRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	var router *gin.Engine
	if conf.ServerConfig.Mode == "dev" {
		router = gin.Default()
	} else {
		router = gin.New()
	}
	if len(middlewares) > 0 {
		router.Use(gin.RecoveryWithWriter(logger.Output()))
		router.Use(middlewares...)
	}

	return router
}
