package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeUserModel = (*customBladeUserModel)(nil)

type (
	// BladeUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeUserModel.
	BladeUserModel interface {
		bladeUserModel
	}

	customBladeUserModel struct {
		*defaultBladeUserModel
	}
)

// NewBladeUserModel returns a model for the database table.
func NewBladeUserModel(conn sqlx.SqlConn, c cache.CacheConf) BladeUserModel {
	return &customBladeUserModel{
		defaultBladeUserModel: newBladeUserModel(conn, c),
	}
}
