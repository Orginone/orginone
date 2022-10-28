package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuTaskModel = (*customActRuTaskModel)(nil)

type (
	// ActRuTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuTaskModel.
	ActRuTaskModel interface {
		actRuTaskModel
	}

	customActRuTaskModel struct {
		*defaultActRuTaskModel
	}
)

// NewActRuTaskModel returns a model for the database table.
func NewActRuTaskModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuTaskModel {
	return &customActRuTaskModel{
		defaultActRuTaskModel: newActRuTaskModel(conn, c),
	}
}
