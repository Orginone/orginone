package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppRoleModel = (*customAsMarketAppRoleModel)(nil)

type (
	// AsMarketAppRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppRoleModel.
	AsMarketAppRoleModel interface {
		asMarketAppRoleModel
	}

	customAsMarketAppRoleModel struct {
		*defaultAsMarketAppRoleModel
	}
)

// NewAsMarketAppRoleModel returns a model for the database table.
func NewAsMarketAppRoleModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppRoleModel {
	return &customAsMarketAppRoleModel{
		defaultAsMarketAppRoleModel: newAsMarketAppRoleModel(conn, c),
	}
}
