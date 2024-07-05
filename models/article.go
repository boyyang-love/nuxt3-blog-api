package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type BaseArticle struct {
	Base
	// 字段
	Title        string `json:"title" form:"title"`
	Des          string `json:"des" form:"des"`
	Cover        string `json:"cover" form:"cover"`
	Content      string `json:"content" form:"content" gorm:"size:15000"`
	Keywords     string `json:"keywords" form:"keywords"`
	Star         int    `json:"star" form:"star"`
	UserId       uint   `json:"user_id" form:"user_id"`
	CategoriesId uint   `json:"categories_id" form:"categories_id"`
}

type Article struct {
	BaseArticle
	// 关系
	User       User       `json:"user" form:"user" gorm:"reference:UserId"`
	Categories Categories `json:"categories" gorm:"reference:CategoriesId"`
	Tag        []*Tag     `json:"tag" form:"tag" gorm:"many2many:article_tag"`
	Comment    []Comment  `json:"comment" form:"comment" gorm:"reference:Id"`
}

func (a *BaseArticle) TableName() string {
	return "article"
}

func (a *Article) TableName() string {
	return "article"
}

func (a *BaseArticle) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	a.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
