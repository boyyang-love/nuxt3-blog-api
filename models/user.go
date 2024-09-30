package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	Base
	// 字段
	Username string `json:"username" form:"username"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	Motto    string `json:"motto" form:"motto" gorm:"default:第一行没有你，第二行没有你，第三行没有也罢！-胡歌"`
	Address  string `json:"address" form:"address"`
	Tel      string `json:"tel" form:"tel"`
	Email    string `json:"email" form:"email"`
	QQ       string `json:"qq" form:"qq"`
	Wechat   string `json:"wechat" form:"wechat"`
	GitHub   string `json:"git_hub" form:"git_hub" gorm:"xxxx@gmail.com"`
	Role     string `json:"role" form:"role" gorm:"default:user"`
	Avatar   string `json:"avatar" form:"avatar" gorm:"default:BOYYANG/default/default_avatar.jpg"`
	Cover    string `json:"cover" form:"cover" gorm:"default:BOYYANG/default/default_cover.jpg"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	u.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
