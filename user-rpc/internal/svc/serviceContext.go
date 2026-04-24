package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user-rpc/internal/config"
	model2 "go-zero-demo/user-rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model2.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model2.NewUserModel(conn),
	}
}
