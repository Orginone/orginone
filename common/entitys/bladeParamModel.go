package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeParamModel = (*customBladeParamModel)(nil)

type (
	// BladeParamModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeParamModel.
	BladeParamModel interface {
		bladeParamModel
	}

	customBladeParamModel struct {
		*defaultBladeParamModel
	}
)

// NewBladeParamModel returns a model for the database table.
func NewBladeParamModel(conn sqlx.SqlConn, c cache.CacheConf) BladeParamModel {
	return &customBladeParamModel{
		defaultBladeParamModel: newBladeParamModel(conn, c),
	}
}
