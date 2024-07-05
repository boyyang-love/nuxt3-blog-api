package config

import "github.com/zeromicro/go-zero/rest"

type MySQLConf struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	Collation string
	Timeout   string
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type CloudBase struct {
	ClientUrl       string
	ClientSecretId  string
	ClientSecretKey string
}

type MinioClient struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Secure    bool
}
type Config struct {
	rest.RestConf
	MySQLConf   MySQLConf
	Auth        Auth
	CloudBase   CloudBase
	MinioClient MinioClient
}
