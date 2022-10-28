package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDeModelModel = (*customActDeModelModel)(nil)

type (
	// ActDeModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDeModelModel.
	ActDeModelModel interface {
		actDeModelModel
	}

	customActDeModelModel struct {
		*defaultActDeModelModel
	}
)

// NewActDeModelModel returns a model for the database table.
func NewActDeModelModel(conn sqlx.SqlConn, c cache.CacheConf) ActDeModelModel {
	return &customActDeModelModel{
		defaultActDeModelModel: newActDeModelModel(conn, c),
	}
}
