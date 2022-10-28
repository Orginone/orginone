package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeRoleModel = (*customBladeRoleModel)(nil)

type (
	// BladeRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeRoleModel.
	BladeRoleModel interface {
		bladeRoleModel
	}

	customBladeRoleModel struct {
		*defaultBladeRoleModel
	}
)

// NewBladeRoleModel returns a model for the database table.
func NewBladeRoleModel(conn sqlx.SqlConn, c cache.CacheConf) BladeRoleModel {
	return &customBladeRoleModel{
		defaultBladeRoleModel: newBladeRoleModel(conn, c),
	}
}
