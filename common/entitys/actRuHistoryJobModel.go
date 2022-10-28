package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuHistoryJobModel = (*customActRuHistoryJobModel)(nil)

type (
	// ActRuHistoryJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuHistoryJobModel.
	ActRuHistoryJobModel interface {
		actRuHistoryJobModel
	}

	customActRuHistoryJobModel struct {
		*defaultActRuHistoryJobModel
	}
)

// NewActRuHistoryJobModel returns a model for the database table.
func NewActRuHistoryJobModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuHistoryJobModel {
	return &customActRuHistoryJobModel{
		defaultActRuHistoryJobModel: newActRuHistoryJobModel(conn, c),
	}
}
