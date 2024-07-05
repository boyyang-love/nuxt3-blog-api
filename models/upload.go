package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type Upload struct {
	Base
	Hash     string `json:"hash"`
	FileName string `json:"file_name" form:"file_name"`
	FileSize int64  `json:"file_size" form:"file_size"`
	FileType string `json:"file_type" form:"file_type"`
	FilePath string `json:"file_path" form:"file_path"`
	UserId   uint   `json:"user_id" form:"user_id"`
	Type     string `json:"type" form:"type"`
	Status   bool   `json:"status" form:"status" gorm:"default:false"`
	W        int    `json:"w" form:"w"`
	H        int    `json:"h" form:"h"`
}

func (u *Upload) TableName() string {
	return "upload"
}

func (u *Upload) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	u.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
