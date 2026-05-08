package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	model2 "go-zero-mall/app/order/service/model"
	"go-zero-mall/app/order/service/rpc/internal/config"
	"go-zero-mall/app/travel/service/model"
)

type ServiceContext struct {
	Config             config.Config
	HomestayModel      model.HomestayModel
	HomestayOrderModel model2.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:             c,
		HomestayModel:      model.NewHomestayModel(sqlConn, c.Cache),
		HomestayOrderModel: model2.NewHomestayOrderModel(sqlConn, c.Cache),
	}
}
