package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuExecutionModel = (*customActRuExecutionModel)(nil)

type (
	// ActRuExecutionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuExecutionModel.
	ActRuExecutionModel interface {
		actRuExecutionModel
	}

	customActRuExecutionModel struct {
		*defaultActRuExecutionModel
	}
)

// NewActRuExecutionModel returns a model for the database table.
func NewActRuExecutionModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuExecutionModel {
	return &customActRuExecutionModel{
		defaultActRuExecutionModel: newActRuExecutionModel(conn, c),
	}
}
