package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsGroupTenantRelationsModel = (*customAsGroupTenantRelationsModel)(nil)

type (
	// AsGroupTenantRelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsGroupTenantRelationsModel.
	AsGroupTenantRelationsModel interface {
		asGroupTenantRelationsModel
	}

	customAsGroupTenantRelationsModel struct {
		*defaultAsGroupTenantRelationsModel
	}
)

// NewAsGroupTenantRelationsModel returns a model for the database table.
func NewAsGroupTenantRelationsModel(conn sqlx.SqlConn, c cache.CacheConf) AsGroupTenantRelationsModel {
	return &customAsGroupTenantRelationsModel{
		defaultAsGroupTenantRelationsModel: newAsGroupTenantRelationsModel(conn, c),
	}
}
