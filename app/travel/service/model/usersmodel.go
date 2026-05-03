package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

func (m *defaultUsersModel) FindColumnList(
	ctx context.Context,
	builder squirrel.SelectBuilder,
	column string,
	orderBy string,
) ([]string, error) {

	queryBuilder := builder.Columns(column)

	if orderBy != "" {
		queryBuilder = queryBuilder.OrderBy(orderBy)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []string

	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
