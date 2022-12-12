package service

import (
	"errors"
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/pkg/code"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils/encrypt"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		dao: dao.NewUserDao(),
	}
}

var ErrUserDeleted = errors.New("user deleted")

func (u *UserService) Login(username, password string) (*model.User, error) {
	user, err := u.dao.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	if code.UserStatus(user.Status) == code.StatusDeleted {
		return nil, ErrUserDeleted
	}

	token, err := encrypt.CreateToken(user.Id, user.Username, user.Uid)
	if err != nil {
		return nil, err
	}
	user.Token = token

	return user, nil
}
