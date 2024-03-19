package config

import "github.com/zeromicro/go-zero/rest"

type MySQLConf struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Charset  string
	Timeout  string
}

type Config struct {
	rest.RestConf
	MySQLConf MySQLConf
}
