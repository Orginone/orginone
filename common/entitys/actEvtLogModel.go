package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActEvtLogModel = (*customActEvtLogModel)(nil)

type (
	// ActEvtLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActEvtLogModel.
	ActEvtLogModel interface {
		actEvtLogModel
	}

	customActEvtLogModel struct {
		*defaultActEvtLogModel
	}
)

// NewActEvtLogModel returns a model for the database table.
func NewActEvtLogModel(conn sqlx.SqlConn, c cache.CacheConf) ActEvtLogModel {
	return &customActEvtLogModel{
		defaultActEvtLogModel: newActEvtLogModel(conn, c),
	}
}
