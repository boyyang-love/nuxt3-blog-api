package svc

import (
	"blog_backend/common/helper"
	"blog_backend/internal/config"
	"fmt"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := helper.ConMySQL(c.MySQLConf)
	if err != nil {
		fmt.Println("mysql连接失败", err.Error())
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
