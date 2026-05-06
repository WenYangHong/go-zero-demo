package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HomestayCommentModel = (*CustomHomestayCommentModel)(nil)

type (
	// HomestayCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayCommentModel.
	HomestayCommentModel interface {
		homestayCommentModel
		FindStatsByHomestayIds(ctx context.Context, homestayIds []int64) ([]*HomestayCommentStats, error)
	}

	CustomHomestayCommentModel struct {
		*defaultHomestayCommentModel
	}
)

// NewHomestayCommentModel returns a model for the database table.
func NewHomestayCommentModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayCommentModel {
	return &CustomHomestayCommentModel{
		defaultHomestayCommentModel: newHomestayCommentModel(conn, c),
	}
}

// model/homestaycomment.go (扩展部分)
type HomestayCommentStats struct {
	HomestayId int64   `db:"homestay_id"`
	Count      int64   `db:"count"`
	AvgScore   float64 `db:"avg_score"`
}

func (m *CustomHomestayCommentModel) FindStatsByHomestayIds(ctx context.Context, homestayIds []int64) ([]*HomestayCommentStats, error) {
	if len(homestayIds) == 0 {
		return []*HomestayCommentStats{}, nil
	}

	builder := m.SelectBuilder().
		Columns(
			"homestay_id",
			"COUNT(id) as count",
			"COALESCE(AVG(avg_star), 0) as avg_score", // 避免无评论时返回 NULL
		).
		Where(squirrel.Eq{"homestay_id": homestayIds}).
		GroupBy("homestay_id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query failed: %w", err)
	}

	var stats []*HomestayCommentStats
	if err := m.QueryRowsNoCacheCtx(ctx, &stats, query, args...); err != nil {
		return nil, fmt.Errorf("execute query failed: %w", err)
	}
	return stats, nil
}
