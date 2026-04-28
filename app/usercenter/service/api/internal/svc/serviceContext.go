package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall/app/usercenter/service/api/internal/config"
	"go-zero-mall/app/usercenter/service/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
