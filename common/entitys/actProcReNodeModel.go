package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActProcReNodeModel = (*customActProcReNodeModel)(nil)

type (
	// ActProcReNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActProcReNodeModel.
	ActProcReNodeModel interface {
		actProcReNodeModel
	}

	customActProcReNodeModel struct {
		*defaultActProcReNodeModel
	}
)

// NewActProcReNodeModel returns a model for the database table.
func NewActProcReNodeModel(conn sqlx.SqlConn, c cache.CacheConf) ActProcReNodeModel {
	return &customActProcReNodeModel{
		defaultActProcReNodeModel: newActProcReNodeModel(conn, c),
	}
}
