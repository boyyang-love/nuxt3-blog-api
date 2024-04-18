package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Categories struct {
	Id        uint       `json:"id" form:"id" gorm:"primaryKey"`
	Uid       string     `json:"uid" form:"uid"`
	Created   int64      `gorm:"autoCreateTime:milli"`
	Updated   int64      `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at"`
	// 字段
	Name   string `json:"name" form:"name"`
	Cover  string `json:"cover" form:"cover"`
	Des    string `json:"des" form:"des"`
	UserId uint   `json:"user_id" form:"user_id"`
}

func (c *Categories) TableName() string {
	return "categories"
}

func (c *Categories) BeforeCreate(*gorm.DB) error {
	uid := uuid.NewV1()
	c.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
