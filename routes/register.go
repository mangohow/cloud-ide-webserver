package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/controller"
	"github.com/mangohow/cloud-ide-webserver/internal/middleware"
)

func Register(engine *gin.Engine) {
	authGroup := engine.Group("/auth")
	userController := controller.NewUserController()
	{
		authGroup.POST("/login", Decorate(userController.Login))
		authGroup.GET("/username/check", Decorate(userController.CheckUsernameAvailable))
		authGroup.POST("/register", Decorate(userController.Register))
		authGroup.GET("/emailCode", Decorate(userController.GetEmailValidateCode))
	}

	apiGroup := engine.Group("/api", middleware.Auth())
	tmplController := controller.NewSpaceTmplController()
	{
		apiGroup.GET("/template/list", Decorate(tmplController.SpaceTmpls))
		apiGroup.GET("/spec/list", Decorate(tmplController.SpaceSpecs))
	}

	spaceController := controller.NewCloudCodeController()
	{
		apiGroup.GET("/workspace/list", Decorate(spaceController.ListSpace))
		apiGroup.DELETE("/workspace", Decorate(spaceController.DeleteSpace))
		apiGroup.POST("/workspace", Decorate(spaceController.CreateSpace))
		apiGroup.POST("/workspace/cas", Decorate(spaceController.CreateSpaceAndStart))
		apiGroup.PUT("/workspace/start", Decorate(spaceController.StartSpace))
		apiGroup.PUT("/workspace/stop", Decorate(spaceController.StopSpace))
		apiGroup.PUT("/workspace/name", Decorate(spaceController.ModifySpaceName))
	}
}
