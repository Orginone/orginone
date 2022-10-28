package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTenantAttrRoleModel = (*customAsTenantAttrRoleModel)(nil)

type (
	// AsTenantAttrRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTenantAttrRoleModel.
	AsTenantAttrRoleModel interface {
		asTenantAttrRoleModel
	}

	customAsTenantAttrRoleModel struct {
		*defaultAsTenantAttrRoleModel
	}
)

// NewAsTenantAttrRoleModel returns a model for the database table.
func NewAsTenantAttrRoleModel(conn sqlx.SqlConn, c cache.CacheConf) AsTenantAttrRoleModel {
	return &customAsTenantAttrRoleModel{
		defaultAsTenantAttrRoleModel: newAsTenantAttrRoleModel(conn, c),
	}
}
