package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketRoleMenuModel = (*customAsMarketRoleMenuModel)(nil)

type (
	// AsMarketRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketRoleMenuModel.
	AsMarketRoleMenuModel interface {
		asMarketRoleMenuModel
	}

	customAsMarketRoleMenuModel struct {
		*defaultAsMarketRoleMenuModel
	}
)

// NewAsMarketRoleMenuModel returns a model for the database table.
func NewAsMarketRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketRoleMenuModel {
	return &customAsMarketRoleMenuModel{
		defaultAsMarketRoleMenuModel: newAsMarketRoleMenuModel(conn, c),
	}
}
