package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeLogApiModel = (*customBladeLogApiModel)(nil)

type (
	// BladeLogApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeLogApiModel.
	BladeLogApiModel interface {
		bladeLogApiModel
	}

	customBladeLogApiModel struct {
		*defaultBladeLogApiModel
	}
)

// NewBladeLogApiModel returns a model for the database table.
func NewBladeLogApiModel(conn sqlx.SqlConn, c cache.CacheConf) BladeLogApiModel {
	return &customBladeLogApiModel{
		defaultBladeLogApiModel: newBladeLogApiModel(conn, c),
	}
}
