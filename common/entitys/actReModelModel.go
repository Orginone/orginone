package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActReModelModel = (*customActReModelModel)(nil)

type (
	// ActReModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActReModelModel.
	ActReModelModel interface {
		actReModelModel
	}

	customActReModelModel struct {
		*defaultActReModelModel
	}
)

// NewActReModelModel returns a model for the database table.
func NewActReModelModel(conn sqlx.SqlConn, c cache.CacheConf) ActReModelModel {
	return &customActReModelModel{
		defaultActReModelModel: newActReModelModel(conn, c),
	}
}
