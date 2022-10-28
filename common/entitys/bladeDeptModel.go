package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeDeptModel = (*customBladeDeptModel)(nil)

type (
	// BladeDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeDeptModel.
	BladeDeptModel interface {
		bladeDeptModel
	}

	customBladeDeptModel struct {
		*defaultBladeDeptModel
	}
)

// NewBladeDeptModel returns a model for the database table.
func NewBladeDeptModel(conn sqlx.SqlConn, c cache.CacheConf) BladeDeptModel {
	return &customBladeDeptModel{
		defaultBladeDeptModel: newBladeDeptModel(conn, c),
	}
}
