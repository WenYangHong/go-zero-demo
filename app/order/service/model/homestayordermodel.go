package model

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HomestayOrderModel = (*customHomestayOrderModel)(nil)

type (
	// HomestayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayOrderModel.
	HomestayOrderModel interface {
		homestayOrderModel
		FindOneByHomestayIdLiveTimeTradeState(ctx context.Context, userId int64, homestayId int64, liveStartTime int64, liveEndTime int64, orderStatus []int64) (bool, error)
	}

	customHomestayOrderModel struct {
		*defaultHomestayOrderModel
	}
)

// NewHomestayOrderModel returns a model for the database table.
func NewHomestayOrderModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayOrderModel {
	return &customHomestayOrderModel{
		defaultHomestayOrderModel: newHomestayOrderModel(conn, c),
	}
}

func (m *customHomestayOrderModel) FindOneByHomestayIdLiveTimeTradeState(ctx context.Context, userId int64, homestayId int64, liveStartTime int64, liveEndTime int64, orderStatus []int64) (bool, error) {
	// 构建查询，使用 COUNT(*) 获取符合条件的数据条数
	builder := m.SelectBuilder().
		Columns("COUNT(*) as count").
		Where(squirrel.Eq{"homestay_id": homestayId}).
		Where(squirrel.Eq{"trade_state": orderStatus}).
		Where(squirrel.GtOrEq{"live_start_date": time.Unix(liveStartTime, 0).Format("2006-01-02")}).
		Where(squirrel.LtOrEq{"live_end_date": time.Unix(liveEndTime, 0).Format("2006-01-02")}).
		Where(squirrel.Eq{"user_id": userId})

	query, args, err := builder.ToSql()
	fmt.Println("-- query", query)
	if err != nil {
		// 构建 SQL 出错，返回 false 和错误信息
		return false, fmt.Errorf("build query failed: %w", err)
	}

	var count int64
	// 执行查询并扫描结果
	if err := m.QueryRowNoCacheCtx(ctx, &count, query, args...); err != nil {
		// 查询执行出错，返回 false 和错误信息
		return false, fmt.Errorf("execute query failed: %w", err)
	}

	// 如果 count > 0 表示存在符合条件的订单，返回 true；否则返回 false
	return count > 0, nil
}
