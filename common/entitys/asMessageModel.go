package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMessageModel = (*customAsMessageModel)(nil)

type (
	// AsMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMessageModel.
	AsMessageModel interface {
		asMessageModel
	}

	customAsMessageModel struct {
		*defaultAsMessageModel
	}
)

// NewAsMessageModel returns a model for the database table.
func NewAsMessageModel(conn sqlx.SqlConn, c cache.CacheConf) AsMessageModel {
	return &customAsMessageModel{
		defaultAsMessageModel: newAsMessageModel(conn, c),
	}
}
