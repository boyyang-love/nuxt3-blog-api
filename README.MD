# boyyang-blog 后端接口

### 项目基本配置（etc/blog-backend.yaml）

```yaml

Name: blog_backend // 项目名称
Host: 0.0.0.0 // 项目启动地址
Port: 9527 // 项目启动端口号
Log:
  Mode: file // 日志打印方式
Auth: 
  AccessSecret: boyyangBlog // token密钥
  AccessExpire: 86_400_000 // token过期时长
MySQLConf: 
  Host: '' // 数据库host
  Port: 9527 // 数据库端口号
  Database: 名称 // 数据库名称
  Username: root // 用户名
  Password: xxxx // 密码
  Charset: utf8mb4 
  Collation: utf8mb4_unicode_ci
  Timeout: 10s
CloudBase:
  ClientUrl: '' // 腾讯云对象存储(现在没有用)
  ClientSecretId: ''
  ClientSecretKey: ''
MinioClient:
  Endpoint: '' // minio host:端口 'host':9000
  AccessKey: '' // access key
  SecretKey: '' // secret key
  Secure: false

```

### 启动

`bash
go run blog-backend.go
`
或者通过安装fresh 或者Air来启动项目

### 打包部署

* 通过docker部署
```dockerfile
FROM golang:1.22.1 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

#RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/etc  etc

EXPOSE 9527

CMD ["/app/main"]

```

* 直接打包二进制文件

```bash
GO111MODULE=on GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
```
