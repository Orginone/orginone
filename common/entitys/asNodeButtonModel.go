package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsNodeButtonModel = (*customAsNodeButtonModel)(nil)

type (
	// AsNodeButtonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsNodeButtonModel.
	AsNodeButtonModel interface {
		asNodeButtonModel
	}

	customAsNodeButtonModel struct {
		*defaultAsNodeButtonModel
	}
)

// NewAsNodeButtonModel returns a model for the database table.
func NewAsNodeButtonModel(conn sqlx.SqlConn, c cache.CacheConf) AsNodeButtonModel {
	return &customAsNodeButtonModel{
		defaultAsNodeButtonModel: newAsNodeButtonModel(conn, c),
	}
}
