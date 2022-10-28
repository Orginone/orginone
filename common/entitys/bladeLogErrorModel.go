package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeLogErrorModel = (*customBladeLogErrorModel)(nil)

type (
	// BladeLogErrorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeLogErrorModel.
	BladeLogErrorModel interface {
		bladeLogErrorModel
	}

	customBladeLogErrorModel struct {
		*defaultBladeLogErrorModel
	}
)

// NewBladeLogErrorModel returns a model for the database table.
func NewBladeLogErrorModel(conn sqlx.SqlConn, c cache.CacheConf) BladeLogErrorModel {
	return &customBladeLogErrorModel{
		defaultBladeLogErrorModel: newBladeLogErrorModel(conn, c),
	}
}
