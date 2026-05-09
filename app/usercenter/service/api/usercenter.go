package main

import (
	"flag"
	"fmt"
	"go-zero-mall/pkg/xerr"
	"net/http"

	"go-zero-mall/app/usercenter/service/api/internal/config"
	"go-zero-mall/app/usercenter/service/api/internal/handler"
	"go-zero-mall/app/usercenter/service/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJson(w, http.StatusNotFound, map[string]any{
			"code": xerr.NOT_FOUND,
			"msg":  "路由不存在",
			"data": nil,
		})
	})))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
