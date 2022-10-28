package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiActinstModel = (*customActHiActinstModel)(nil)

type (
	// ActHiActinstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiActinstModel.
	ActHiActinstModel interface {
		actHiActinstModel
	}

	customActHiActinstModel struct {
		*defaultActHiActinstModel
	}
)

// NewActHiActinstModel returns a model for the database table.
func NewActHiActinstModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiActinstModel {
	return &customActHiActinstModel{
		defaultActHiActinstModel: newActHiActinstModel(conn, c),
	}
}
