package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeRoleMenuModel = (*customBladeRoleMenuModel)(nil)

type (
	// BladeRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeRoleMenuModel.
	BladeRoleMenuModel interface {
		bladeRoleMenuModel
	}

	customBladeRoleMenuModel struct {
		*defaultBladeRoleMenuModel
	}
)

// NewBladeRoleMenuModel returns a model for the database table.
func NewBladeRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf) BladeRoleMenuModel {
	return &customBladeRoleMenuModel{
		defaultBladeRoleMenuModel: newBladeRoleMenuModel(conn, c),
	}
}
