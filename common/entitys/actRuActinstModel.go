package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuActinstModel = (*customActRuActinstModel)(nil)

type (
	// ActRuActinstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuActinstModel.
	ActRuActinstModel interface {
		actRuActinstModel
	}

	customActRuActinstModel struct {
		*defaultActRuActinstModel
	}
)

// NewActRuActinstModel returns a model for the database table.
func NewActRuActinstModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuActinstModel {
	return &customActRuActinstModel{
		defaultActRuActinstModel: newActRuActinstModel(conn, c),
	}
}
