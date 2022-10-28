package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsProcValueModel = (*customAsProcValueModel)(nil)

type (
	// AsProcValueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsProcValueModel.
	AsProcValueModel interface {
		asProcValueModel
	}

	customAsProcValueModel struct {
		*defaultAsProcValueModel
	}
)

// NewAsProcValueModel returns a model for the database table.
func NewAsProcValueModel(conn sqlx.SqlConn, c cache.CacheConf) AsProcValueModel {
	return &customAsProcValueModel{
		defaultAsProcValueModel: newAsProcValueModel(conn, c),
	}
}
