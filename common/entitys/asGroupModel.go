package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsGroupModel = (*customAsGroupModel)(nil)

type (
	// AsGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsGroupModel.
	AsGroupModel interface {
		asGroupModel
	}

	customAsGroupModel struct {
		*defaultAsGroupModel
	}
)

// NewAsGroupModel returns a model for the database table.
func NewAsGroupModel(conn sqlx.SqlConn, c cache.CacheConf) AsGroupModel {
	return &customAsGroupModel{
		defaultAsGroupModel: newAsGroupModel(conn, c),
	}
}
