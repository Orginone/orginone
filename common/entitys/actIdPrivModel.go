package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdPrivModel = (*customActIdPrivModel)(nil)

type (
	// ActIdPrivModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdPrivModel.
	ActIdPrivModel interface {
		actIdPrivModel
	}

	customActIdPrivModel struct {
		*defaultActIdPrivModel
	}
)

// NewActIdPrivModel returns a model for the database table.
func NewActIdPrivModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdPrivModel {
	return &customActIdPrivModel{
		defaultActIdPrivModel: newActIdPrivModel(conn, c),
	}
}
