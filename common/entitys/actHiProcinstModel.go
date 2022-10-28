package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiProcinstModel = (*customActHiProcinstModel)(nil)

type (
	// ActHiProcinstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiProcinstModel.
	ActHiProcinstModel interface {
		actHiProcinstModel
	}

	customActHiProcinstModel struct {
		*defaultActHiProcinstModel
	}
)

// NewActHiProcinstModel returns a model for the database table.
func NewActHiProcinstModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiProcinstModel {
	return &customActHiProcinstModel{
		defaultActHiProcinstModel: newActHiProcinstModel(conn, c),
	}
}
