package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-mall/app/travel/service/api/internal/config"
	"go-zero-mall/app/travel/service/model"
)

type ServiceContext struct {
	Config                config.Config
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayCategoryModel model.HomestayCategoryModel
	CategoryModel         model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		HomestayModel:         model.NewHomestayModel(sqlConn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlConn, c.Cache),
		HomestayCategoryModel: model.NewHomestayCategoryModel(sqlConn, c.Cache),
		CategoryModel:         model.NewCategoryModel(sqlConn, c.Cache),
	}
}
