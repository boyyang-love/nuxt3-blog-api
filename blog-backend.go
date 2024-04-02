package main

import (
	"blog_backend/common/errorx"
	"blog_backend/common/respx"
	"blog_backend/internal/config"
	"blog_backend/internal/handler"
	"blog_backend/internal/svc"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog-backend.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCustomCors(
			nil,
			notAllowedFn,
			"http://www.boyyang.cn",
			"http://localhost:3000",
		),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 设置返回结果
	httpx.SetOkHandler(func(ctx context.Context, data interface{}) (r interface{}) {

		return &respx.Body{
			Msg:  "ok",
			Code: 1,
			Data: data,
		}
	})
	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		var e *errorx.CodeError
		switch {
		case errors.As(err, &e):
			return http.StatusInternalServerError, e.Data()
		default:
			e.Code = 0
			e.Msg = err.Error()
			return http.StatusInternalServerError, e.Data()
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
