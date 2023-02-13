package service

import (
	"errors"
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils/encrypt"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	logger *logrus.Logger
	dao    *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		logger: logger.Logger(),
		dao:    dao.NewUserDao(),
	}
}

var ErrUserDeleted = errors.New("user deleted")

func (u *UserService) Login(username, password string) (*model.User, error) {
	// 1、从数据库中查询
	user, err := u.dao.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	// 2、检查用户状态是否正常
	if code.UserStatus(user.Status) == code.StatusDeleted {
		return nil, ErrUserDeleted
	}

	// 3、生成token
	token, err := encrypt.CreateToken(user.Id, user.Username, user.Uid)
	if err != nil {
		return nil, err
	}
	user.Token = token

	return user, nil
}

func (u *UserService) CheckUsernameAvailable(username string) bool {
	err := u.dao.FindByUsername(username)
	// 如果能查询到记录， err == nil
	if err != nil {
		return false
	}

	return true
}
