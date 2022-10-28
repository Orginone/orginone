package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeLogUsualModel = (*customBladeLogUsualModel)(nil)

type (
	// BladeLogUsualModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeLogUsualModel.
	BladeLogUsualModel interface {
		bladeLogUsualModel
	}

	customBladeLogUsualModel struct {
		*defaultBladeLogUsualModel
	}
)

// NewBladeLogUsualModel returns a model for the database table.
func NewBladeLogUsualModel(conn sqlx.SqlConn, c cache.CacheConf) BladeLogUsualModel {
	return &customBladeLogUsualModel{
		defaultBladeLogUsualModel: newBladeLogUsualModel(conn, c),
	}
}
