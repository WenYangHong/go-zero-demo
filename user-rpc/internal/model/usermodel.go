package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		CountActiveUsers(ctx context.Context) (int64, error)
		TransactCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultUserModel) CountActiveUsers(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE `created_at` > '2026-04-24 10:00:00'", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query)
	return count, err
}

// TransactCtx 执行一个带上下文的数据库事务
func (m *defaultUserModel) TransactCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	// 直接调用底层 sqlx.SqlConn 的事务方法
	return m.conn.TransactCtx(ctx, fn)
}
