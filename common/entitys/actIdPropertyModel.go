package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdPropertyModel = (*customActIdPropertyModel)(nil)

type (
	// ActIdPropertyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdPropertyModel.
	ActIdPropertyModel interface {
		actIdPropertyModel
	}

	customActIdPropertyModel struct {
		*defaultActIdPropertyModel
	}
)

// NewActIdPropertyModel returns a model for the database table.
func NewActIdPropertyModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdPropertyModel {
	return &customActIdPropertyModel{
		defaultActIdPropertyModel: newActIdPropertyModel(conn, c),
	}
}
