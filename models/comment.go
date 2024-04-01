package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type BaseComment struct {
	Id        uint       `json:"id" form:"id" gorm:"primaryKey"`
	Uid       string     `json:"uid" form:"uid"`
	Created   int64      `gorm:"autoCreateTime:milli"`
	Updated   int64      `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at"`
	// 字段
	Content       string `json:"content" form:"content"`
	ArticleId     uint   `json:"article_id" form:"article_id"`
	CommentId     uint   `json:"comment_id" form:"comment_id"`
	WebsiteUserId uint   `json:"website_user_id" form:"website_user_id"`
	UserId        uint   `json:"user_id" form:"user_id"`
	Type          string `json:"type" form:"type" gorm:"type:enum('article','comment','website')"`
}

type Comment struct {
	BaseComment
	User User `json:"user" form:"user" gorm:"reference:UserId"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *BaseComment) TableName() string {
	return "comment"
}

func (c *BaseComment) BeforeCreate(*gorm.DB) error {
	uid := uuid.NewV1()
	c.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
