package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeClientModel = (*customBladeClientModel)(nil)

type (
	// BladeClientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeClientModel.
	BladeClientModel interface {
		bladeClientModel
	}

	customBladeClientModel struct {
		*defaultBladeClientModel
	}
)

// NewBladeClientModel returns a model for the database table.
func NewBladeClientModel(conn sqlx.SqlConn, c cache.CacheConf) BladeClientModel {
	return &customBladeClientModel{
		defaultBladeClientModel: newBladeClientModel(conn, c),
	}
}
