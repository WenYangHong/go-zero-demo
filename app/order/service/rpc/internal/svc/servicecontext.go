package svc

import (
	model2 "go-zero-mall/app/order/service/model"
	"go-zero-mall/app/order/service/rpc/internal/config"
	"go-zero-mall/app/travel/service/model"
	"go-zero-mall/pkg/mq"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	HomestayModel      model.HomestayModel
	HomestayOrderModel model2.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	mq.InitRabbitMQ()
	err := mq.InitOrderDelayQueue()
	if err != nil {
		logx.Errorf("InitOrderDelayQueue failed: %v, MQ related features will be disabled", err)
	}
	HomestayOrderModel := model2.NewHomestayOrderModel(sqlConn, c.Cache)
	if err == nil {
		mq.StartOrderConsumer(HomestayOrderModel)
	}
	return &ServiceContext{
		Config:             c,
		HomestayModel:      model.NewHomestayModel(sqlConn, c.Cache),
		HomestayOrderModel: HomestayOrderModel,
	}
}
