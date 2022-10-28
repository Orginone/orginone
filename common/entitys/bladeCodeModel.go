package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeCodeModel = (*customBladeCodeModel)(nil)

type (
	// BladeCodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeCodeModel.
	BladeCodeModel interface {
		bladeCodeModel
	}

	customBladeCodeModel struct {
		*defaultBladeCodeModel
	}
)

// NewBladeCodeModel returns a model for the database table.
func NewBladeCodeModel(conn sqlx.SqlConn, c cache.CacheConf) BladeCodeModel {
	return &customBladeCodeModel{
		defaultBladeCodeModel: newBladeCodeModel(conn, c),
	}
}
