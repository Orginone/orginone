package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormIconModel = (*customAsFormIconModel)(nil)

type (
	// AsFormIconModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormIconModel.
	AsFormIconModel interface {
		asFormIconModel
	}

	customAsFormIconModel struct {
		*defaultAsFormIconModel
	}
)

// NewAsFormIconModel returns a model for the database table.
func NewAsFormIconModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormIconModel {
	return &customAsFormIconModel{
		defaultAsFormIconModel: newAsFormIconModel(conn, c),
	}
}
