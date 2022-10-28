package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsItopsModel = (*customAsItopsModel)(nil)

type (
	// AsItopsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsItopsModel.
	AsItopsModel interface {
		asItopsModel
	}

	customAsItopsModel struct {
		*defaultAsItopsModel
	}
)

// NewAsItopsModel returns a model for the database table.
func NewAsItopsModel(conn sqlx.SqlConn, c cache.CacheConf) AsItopsModel {
	return &customAsItopsModel{
		defaultAsItopsModel: newAsItopsModel(conn, c),
	}
}
