package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMenuModel = (*customAsMenuModel)(nil)

type (
	// AsMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMenuModel.
	AsMenuModel interface {
		asMenuModel
	}

	customAsMenuModel struct {
		*defaultAsMenuModel
	}
)

// NewAsMenuModel returns a model for the database table.
func NewAsMenuModel(conn sqlx.SqlConn, c cache.CacheConf) AsMenuModel {
	return &customAsMenuModel{
		defaultAsMenuModel: newAsMenuModel(conn, c),
	}
}
