package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuJobModel = (*customActRuJobModel)(nil)

type (
	// ActRuJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuJobModel.
	ActRuJobModel interface {
		actRuJobModel
	}

	customActRuJobModel struct {
		*defaultActRuJobModel
	}
)

// NewActRuJobModel returns a model for the database table.
func NewActRuJobModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuJobModel {
	return &customActRuJobModel{
		defaultActRuJobModel: newActRuJobModel(conn, c),
	}
}
