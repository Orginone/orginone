package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsButtonModel = (*customAsButtonModel)(nil)

type (
	// AsButtonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsButtonModel.
	AsButtonModel interface {
		asButtonModel
	}

	customAsButtonModel struct {
		*defaultAsButtonModel
	}
)

// NewAsButtonModel returns a model for the database table.
func NewAsButtonModel(conn sqlx.SqlConn, c cache.CacheConf) AsButtonModel {
	return &customAsButtonModel{
		defaultAsButtonModel: newAsButtonModel(conn, c),
	}
}
