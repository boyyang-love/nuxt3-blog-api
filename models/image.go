package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Image struct {
	Id        uint       `json:"id" form:"id" gorm:"primaryKey"`
	Uid       string     `json:"uid" form:"uid"`
	Created   int64      `gorm:"autoCreateTime:milli"`
	Updated   int64      `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at"`
	// 字段
	Name   string `json:"title" form:"title"`
	Path   string `json:"path" form:"path"`
	Hash   string `json:"hash" form:"hash"`
	Size   int    `json:"size" form:"size"`
	Type   string `json:"type" form:"type"`
	UserId uint   `json:"user_id" form:"user_id"`
	// 关系
	User User   `json:"user" form:"user" gorm:"reference:UserId"`
	Tag  []*Tag `json:"tag" form:"tag" gorm:"many2many:image_tag;"`
}

func (i *Image) TableName() string {
	return "image"
}

func (i *Image) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	i.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
