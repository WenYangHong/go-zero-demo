package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HomestayCategoryModel = (*customHomestayCategoryModel)(nil)

type (
	// HomestayCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayCategoryModel.
	HomestayCategoryModel interface {
		homestayCategoryModel
	}

	customHomestayCategoryModel struct {
		*defaultHomestayCategoryModel
	}
)

// NewHomestayCategoryModel returns a model for the database table.
func NewHomestayCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayCategoryModel {
	return &customHomestayCategoryModel{
		defaultHomestayCategoryModel: newHomestayCategoryModel(conn, c),
	}
}
