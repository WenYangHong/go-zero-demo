package svc

import (
	"go-zero-demo/user-rpc/internal/config"
	model2 "go-zero-demo/user-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model2.UserModel
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model2.NewUserModel(conn),
		Redis:     redis.MustNewRedis(c.CacheRedis),
	}
}
