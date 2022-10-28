package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiDetailModel = (*customActHiDetailModel)(nil)

type (
	// ActHiDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiDetailModel.
	ActHiDetailModel interface {
		actHiDetailModel
	}

	customActHiDetailModel struct {
		*defaultActHiDetailModel
	}
)

// NewActHiDetailModel returns a model for the database table.
func NewActHiDetailModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiDetailModel {
	return &customActHiDetailModel{
		defaultActHiDetailModel: newActHiDetailModel(conn, c),
	}
}
