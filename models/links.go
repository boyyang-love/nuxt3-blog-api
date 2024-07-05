package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Links struct {
	Base
	WebsiteName string `json:"website_name" gorm:"size:20"`
	WebsiteUrl  string `json:"website_url" gorm:"size:100"`
	WebsiteIcon string `json:"website_icon" gorm:"size:100"`
	Email       string `json:"email" gorm:"size:20"`
	Status      int    `json:"status" gorm:"type:enum('1','2','3','4')"` // 1 审核中 2 审核通过 3 审核失败 4 失联
}

func (l *Links) TableName() string {
	return "links"
}

func (l *Links) BeforeCreate(*gorm.DB) error {
	uid := uuid.NewV1()
	l.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
