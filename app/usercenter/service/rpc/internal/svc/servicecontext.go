package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-mall/app/usercenter/service/rpc/internal/config"
	"go-zero-mall/app/usercenter/service/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlConn, c.Cache),
	}
}
