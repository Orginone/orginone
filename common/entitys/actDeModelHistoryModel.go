package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDeModelHistoryModel = (*customActDeModelHistoryModel)(nil)

type (
	// ActDeModelHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDeModelHistoryModel.
	ActDeModelHistoryModel interface {
		actDeModelHistoryModel
	}

	customActDeModelHistoryModel struct {
		*defaultActDeModelHistoryModel
	}
)

// NewActDeModelHistoryModel returns a model for the database table.
func NewActDeModelHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) ActDeModelHistoryModel {
	return &customActDeModelHistoryModel{
		defaultActDeModelHistoryModel: newActDeModelHistoryModel(conn, c),
	}
}
