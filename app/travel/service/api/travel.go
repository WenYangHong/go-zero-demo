package main

import (
	"flag"
	"fmt"
	"go-zero-mall/pkg/xerr"
	"net/http"

	"go-zero-mall/app/travel/service/api/internal/config"
	"go-zero-mall/app/travel/service/api/internal/handler"
	"go-zero-mall/app/travel/service/api/internal/middleware"
	"go-zero-mall/app/travel/service/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/travel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJson(w, http.StatusNotFound, middleware.Response{
			Code: xerr.NOT_FOUND,
			Msg:  "路由不存在",
			Data: nil,
		})
	})))
	defer server.Stop()

	server.Use(middleware.ResponseMiddleware)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
