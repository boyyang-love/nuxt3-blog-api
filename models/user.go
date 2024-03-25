package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	Id        uint       `json:"id" form:"id" gorm:"primaryKey"`
	Uid       string     `json:"uid" form:"uid"`
	Created   int64      `gorm:"autoCreateTime:milli"`
	Updated   int64      `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at"`
	// 字段
	Username string `json:"username" form:"username"`
	Avatar   string `json:"avatar" form:"avatar"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	Motto    string `json:"motto" form:"motto"`
	Address  string `json:"address" form:"address"`
	Tel      int    `json:"tel" form:"tel"`
	Email    string `json:"email" form:"email"`
	QQ       int    `json:"qq" form:"qq"`
	Wechat   string `json:"wechat" form:"wechat"`
	GitHub   string `json:"git_hub" form:"git_hub"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	u.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
