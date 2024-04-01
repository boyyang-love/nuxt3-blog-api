package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Article struct {
	Id        uint       `json:"id" form:"id" gorm:"primaryKey"`
	Uid       string     `json:"uid" form:"uid"`
	Created   int64      `gorm:"autoCreateTime:milli"`
	Updated   int64      `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at"`
	// 字段
	Title   string `json:"title" form:"title"`
	Des     string `json:"des" form:"des"`
	Cover   string `json:"cover" form:"cover"`
	Content string `json:"content" form:"content" gorm:"size:15000"`
	UserId  uint   `json:"user_id" form:"user_id"`
	Star    int    `json:"star" form:"star"`
	// 关系
	User    User      `json:"user" form:"user" gorm:"reference:UserId"`
	Tag     []*Tag    `json:"tag" form:"tag" gorm:"many2many:article_tag"`
	Comment []Comment `json:"comment" form:"comment" gorm:"reference:Id"`
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	a.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
