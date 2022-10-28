package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActReProcdefModel = (*customActReProcdefModel)(nil)

type (
	// ActReProcdefModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActReProcdefModel.
	ActReProcdefModel interface {
		actReProcdefModel
	}

	customActReProcdefModel struct {
		*defaultActReProcdefModel
	}
)

// NewActReProcdefModel returns a model for the database table.
func NewActReProcdefModel(conn sqlx.SqlConn, c cache.CacheConf) ActReProcdefModel {
	return &customActReProcdefModel{
		defaultActReProcdefModel: newActReProcdefModel(conn, c),
	}
}
