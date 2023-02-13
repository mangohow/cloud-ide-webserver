package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/mangohow/cloud-ide-webserver/internal/dao/db"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
)

type UserDao struct {
	db *sqlx.DB
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: db.DB(),
	}
}

func (u *UserDao) FindByUsernameAndPassword(username, password string) (user *model.User, _ error) {
	sql := `SELECT id, uid, username, nickname, email, avatar, status FROM t_user WHERE username = ? AND password = ?`
	user = &model.User{}
	err := u.db.Get(user, sql, username, password)
	return user, err
}

func (u *UserDao) FindByUsername(username string) error {
	sql := "SELECT 1 FROM t_user WHERE username = ?"
	var n int
	return u.db.Get(&n, sql, username)
}
