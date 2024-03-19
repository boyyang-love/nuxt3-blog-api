package main

import (
	"blog_backend/common/errorx"
	"blog_backend/common/respx"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"blog_backend/internal/config"
	"blog_backend/internal/handler"
	"blog_backend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog-backend.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 设置返回结果
	httpx.SetOkHandler(func(ctx context.Context, data interface{}) (r interface{}) {

		return &respx.Body{
			Msg:  "OK",
			Code: 1,
			Data: data,
		}
	})
	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		var e *errorx.CodeError
		var s *errorx.CodeErrorWithStatus
		switch {
		case errors.As(err, &e):
			return http.StatusOK, e.Data()
		case errors.As(err, &s):
			return http.StatusInternalServerError, s.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
