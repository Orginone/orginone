package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiTskLogModel = (*customActHiTskLogModel)(nil)

type (
	// ActHiTskLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiTskLogModel.
	ActHiTskLogModel interface {
		actHiTskLogModel
	}

	customActHiTskLogModel struct {
		*defaultActHiTskLogModel
	}
)

// NewActHiTskLogModel returns a model for the database table.
func NewActHiTskLogModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiTskLogModel {
	return &customActHiTskLogModel{
		defaultActHiTskLogModel: newActHiTskLogModel(conn, c),
	}
}
