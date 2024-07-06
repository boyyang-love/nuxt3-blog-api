package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Comment struct {
	Base
	// 字段
	Content       string `json:"content" form:"content"`
	ArticleId     uint   `json:"article_id" form:"article_id"`
	CommentId     uint   `json:"comment_id" form:"comment_id"`
	WebsiteUserId uint   `json:"website_user_id" form:"website_user_id"`
	UserId        uint   `json:"user_id" form:"user_id"`
	Type          string `json:"type" form:"type" gorm:"type:enum('article','comment','website')"`
	User          User   `json:"user" form:"user" gorm:"reference:UserId"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) BeforeCreate(*gorm.DB) error {
	uid := uuid.NewV1()
	c.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
