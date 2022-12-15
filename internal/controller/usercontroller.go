package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/service"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/serialize"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	logger  *logrus.Logger
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
		logger:  logger.Logger(),
	}
}

// Login method: POST path: /login
func (u *UserController) Login(ctx *gin.Context) *serialize.Response {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	u.logger.Debugf("username:%s passowrd:%s", username, password)
	if username == "" || password == "" {
		return serialize.NewResponseOk(code.LoginFailed, nil)
	}

	user, err := u.service.Login(username, username)
	if err != nil {
		if err == service.ErrUserDeleted {
			return serialize.NewResponseOk(code.LoginUserDeleted, nil)
		}

		u.logger.Warnf("login error:%v", err)
		return serialize.NewResponseOk(code.LoginFailed, nil)
	}

	return serialize.NewResponseOk(code.LoginSuccess, user)
}
