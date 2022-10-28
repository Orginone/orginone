package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiTaskinstModel = (*customActHiTaskinstModel)(nil)

type (
	// ActHiTaskinstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiTaskinstModel.
	ActHiTaskinstModel interface {
		actHiTaskinstModel
	}

	customActHiTaskinstModel struct {
		*defaultActHiTaskinstModel
	}
)

// NewActHiTaskinstModel returns a model for the database table.
func NewActHiTaskinstModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiTaskinstModel {
	return &customActHiTaskinstModel{
		defaultActHiTaskinstModel: newActHiTaskinstModel(conn, c),
	}
}
