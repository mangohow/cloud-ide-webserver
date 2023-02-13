package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/service"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/serialize"
	"github.com/sirupsen/logrus"
	"net/http"
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

// Register 用户注册 method: POST path: /register
func (u *UserController) Register(ctx *gin.Context) *serialize.Response {

	return serialize.NewResponseOKND(0)
}

// CheckUsernameAvailable 检测用户名是否可用 method: GET path: /uaval
func (u *UserController) CheckUsernameAvailable(ctx *gin.Context) *serialize.Response {
	value := ctx.Query("username")
	if value == "" {
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	ok := u.service.CheckUsernameAvailable(value)
	if !ok {
		return serialize.NewResponseOKND(code.UserNameAvailable)
	}

	return serialize.NewResponseOKND(code.UserNameUnavailable)
}
