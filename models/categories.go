package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Categories struct {
	Base
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
