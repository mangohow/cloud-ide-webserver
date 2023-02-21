package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/internal/service"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/serialize"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserController struct {
	logger       *logrus.Logger
	service      *service.UserService
	emailService *service.EmailService
}

func NewUserController() *UserController {
	emailService := service.NewEmailService()
	err := emailService.Start()
	if err != nil {
		panic(err)
	}
	return &UserController{
		service:      service.NewUserService(),
		logger:       logger.Logger(),
		emailService: emailService,
	}
}

// Login method: POST path: /login
func (u *UserController) Login(ctx *gin.Context) *serialize.Response {
	var userInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		userInfo.Username = ctx.PostForm("username")
		userInfo.Password = ctx.PostForm("password")
	}

	u.logger.Debugf("username:%s passowrd:%s", userInfo.Username, userInfo.Password)
	if userInfo.Username == "" || userInfo.Password == "" {
		return serialize.NewResponseOk(code.LoginFailed, nil)
	}
	u.logger.Debugf("username:%s, pasword:%s", userInfo.Username, userInfo.Password)

	user, err := u.service.Login(userInfo.Username, userInfo.Password)
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
	var info model.RegisterInfo
	err := ctx.ShouldBind(&info)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	u.logger.Debug("register", info)
	// 验证EmailCode长度
	if len(info.EmailCode) != 6 {
		return serialize.NewResponseOKND(code.UserEmailCodeInvalid)
	}

	// 验证邮箱有效性
	if !utils.VerifyEmailFormat(info.Email) {
		return serialize.NewResponseOKND(code.UserEmailInvalid)
	}

	err = u.service.UserRegister(&info)
	switch err {
	case service.ErrEmailCodeIncorrect:
		return serialize.NewResponseOKND(code.UserEmailCodeIncorrect)
	case service.ErrEmailAlreadyInUse:
		return serialize.NewResponseOKND(code.UserEmailAlreadyInUse)
	case nil:
		return serialize.NewResponseOKND(code.UserRegisterSuccess)
	}

	u.logger.Debugf("add user err:%v", err)
	return serialize.NewResponseOKND(code.UserRegisterFailed)
}

// CheckUsernameAvailable 检测用户名是否可用 method: GET path: /uname_available
func (u *UserController) CheckUsernameAvailable(ctx *gin.Context) *serialize.Response {
	u.logger.Debugf("check username available")
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

// GetEmailValidateCode 通过邮箱获取验证码 method: GET path: /validate_code
func (u *UserController) GetEmailValidateCode(ctx *gin.Context) *serialize.Response {
	addr := ctx.Query("email")
	if addr == "" {
		ctx.Status(http.StatusBadRequest)
		return nil
	}

	err := u.emailService.Send(addr)
	if err != nil {
		return serialize.NewResponseOKND(code.UserSendValidateCodeFailed)
	}

	return serialize.NewResponseOKND(code.UserSendValidateCodeSuccess)
}
