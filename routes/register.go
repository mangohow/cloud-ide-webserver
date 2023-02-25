package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/controller"
	"github.com/mangohow/cloud-ide-webserver/internal/middleware"
)

func Register(engine *gin.Engine) {
	userController := controller.NewUserController()
	{
		engine.POST("/login", Decorate(userController.Login))
		engine.GET("/uname_available", Decorate(userController.CheckUsernameAvailable))
		engine.POST("/register", Decorate(userController.Register))
		engine.GET("/validate_code", Decorate(userController.GetEmailValidateCode))
	}
	apiGroup := engine.Group("/api", middleware.Auth())
	tmplController := controller.NewSpaceTmplController()
	{
		apiGroup.GET("/tmpls", Decorate(tmplController.SpaceTmpls))
		apiGroup.GET("/specs", Decorate(tmplController.SpaceSpecs))
	}
	spaceController := controller.NewCloudCodeController()
	{
		apiGroup.GET("/spaces", Decorate(spaceController.ListSpace))
		apiGroup.DELETE("/space", Decorate(spaceController.DeleteSpace))
		apiGroup.POST("/space", Decorate(spaceController.CreateSpace))
		apiGroup.POST("/space_cas", Decorate(spaceController.CreateSpaceAndStart))
		apiGroup.PUT("/space_start", Decorate(spaceController.StartSpace))
		apiGroup.PUT("/space_stop", Decorate(spaceController.StopSpace))
	}
}
