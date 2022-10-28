package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeMenuModel = (*customBladeMenuModel)(nil)

type (
	// BladeMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeMenuModel.
	BladeMenuModel interface {
		bladeMenuModel
	}

	customBladeMenuModel struct {
		*defaultBladeMenuModel
	}
)

// NewBladeMenuModel returns a model for the database table.
func NewBladeMenuModel(conn sqlx.SqlConn, c cache.CacheConf) BladeMenuModel {
	return &customBladeMenuModel{
		defaultBladeMenuModel: newBladeMenuModel(conn, c),
	}
}
