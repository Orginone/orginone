package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeDictModel = (*customBladeDictModel)(nil)

type (
	// BladeDictModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeDictModel.
	BladeDictModel interface {
		bladeDictModel
	}

	customBladeDictModel struct {
		*defaultBladeDictModel
	}
)

// NewBladeDictModel returns a model for the database table.
func NewBladeDictModel(conn sqlx.SqlConn, c cache.CacheConf) BladeDictModel {
	return &customBladeDictModel{
		defaultBladeDictModel: newBladeDictModel(conn, c),
	}
}
