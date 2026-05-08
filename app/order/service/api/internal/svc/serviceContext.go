package svc

import (
	"go-zero-mall/app/order/service/api/internal/config"
	"go-zero-mall/app/order/service/rpc/order"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:   c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
