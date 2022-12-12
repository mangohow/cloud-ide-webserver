package model

import "time"

type User struct {
	Id         uint32    `json:"id" db:"id"`
	Uid        string    `json:"uid" db:"uid"`
	Username   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Nickname   string    `json:"nickname" db:"nickname"`
	Email      string    `json:"email" db:"email"`
	Phone      string    `json:"phone" db:"phone"`
	Avatar     string    `json:"avatar" db:"avatar"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	DeleteTime time.Time `json:"delete_time" db:"delete_time"`
	Status     uint32    `json:"status" db:"status"` // 状态 0正常 1已删除

	Token string `json:"token"`
}
