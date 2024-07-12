package svc

import (
	"blog_backend/common/helper"
	"blog_backend/internal/config"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/minio/minio-go/v7"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	Client      *cos.Client
	Cache       *bigcache.BigCache
	MinIoClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := helper.ConMySQL(c.MySQLConf)
	if err != nil {
		fmt.Println("mysql连接失败", err.Error())
	} else {
		fmt.Println("mysql连接成功")
	}

	//client := helper.InitCloudBase(c.CloudBase.ClientUrl, c.CloudBase.ClientSecretId, c.CloudBase.ClientSecretKey)
	//if client != nil {
	//	fmt.Println("腾讯云初始化成功")
	//}

	cache := helper.NewCache()
	err = cache.Init()
	if err == nil {
		fmt.Println("缓存初始化成功")
	}

	minIoClient, err := helper.Minio(c.MinioClient.Endpoint, c.MinioClient.AccessKey, c.MinioClient.SecretKey, c.MinioClient.Secure)
	if err != nil {
		fmt.Println("minio连接失败", err.Error())
	} else {
		fmt.Println("minio连接成功")
	}

	return &ServiceContext{
		Config:      c,
		DB:          db,
		Client:      nil,
		Cache:       cache.BigCache,
		MinIoClient: minIoClient,
	}
}
