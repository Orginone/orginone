package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuVariableModel = (*customActRuVariableModel)(nil)

type (
	// ActRuVariableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuVariableModel.
	ActRuVariableModel interface {
		actRuVariableModel
	}

	customActRuVariableModel struct {
		*defaultActRuVariableModel
	}
)

// NewActRuVariableModel returns a model for the database table.
func NewActRuVariableModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuVariableModel {
	return &customActRuVariableModel{
		defaultActRuVariableModel: newActRuVariableModel(conn, c),
	}
}
