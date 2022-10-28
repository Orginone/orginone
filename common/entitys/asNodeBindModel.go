package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsNodeBindModel = (*customAsNodeBindModel)(nil)

type (
	// AsNodeBindModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsNodeBindModel.
	AsNodeBindModel interface {
		asNodeBindModel
	}

	customAsNodeBindModel struct {
		*defaultAsNodeBindModel
	}
)

// NewAsNodeBindModel returns a model for the database table.
func NewAsNodeBindModel(conn sqlx.SqlConn, c cache.CacheConf) AsNodeBindModel {
	return &customAsNodeBindModel{
		defaultAsNodeBindModel: newAsNodeBindModel(conn, c),
	}
}
