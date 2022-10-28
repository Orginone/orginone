package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiVarinstModel = (*customActHiVarinstModel)(nil)

type (
	// ActHiVarinstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiVarinstModel.
	ActHiVarinstModel interface {
		actHiVarinstModel
	}

	customActHiVarinstModel struct {
		*defaultActHiVarinstModel
	}
)

// NewActHiVarinstModel returns a model for the database table.
func NewActHiVarinstModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiVarinstModel {
	return &customActHiVarinstModel{
		defaultActHiVarinstModel: newActHiVarinstModel(conn, c),
	}
}
