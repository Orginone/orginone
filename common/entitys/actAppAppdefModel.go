package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAppAppdefModel = (*customActAppAppdefModel)(nil)

type (
	// ActAppAppdefModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAppAppdefModel.
	ActAppAppdefModel interface {
		actAppAppdefModel
	}

	customActAppAppdefModel struct {
		*defaultActAppAppdefModel
	}
)

// NewActAppAppdefModel returns a model for the database table.
func NewActAppAppdefModel(conn sqlx.SqlConn, c cache.CacheConf) ActAppAppdefModel {
	return &customActAppAppdefModel{
		defaultActAppAppdefModel: newActAppAppdefModel(conn, c),
	}
}
