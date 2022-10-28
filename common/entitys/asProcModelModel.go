package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsProcModelModel = (*customAsProcModelModel)(nil)

type (
	// AsProcModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsProcModelModel.
	AsProcModelModel interface {
		asProcModelModel
	}

	customAsProcModelModel struct {
		*defaultAsProcModelModel
	}
)

// NewAsProcModelModel returns a model for the database table.
func NewAsProcModelModel(conn sqlx.SqlConn, c cache.CacheConf) AsProcModelModel {
	return &customAsProcModelModel{
		defaultAsProcModelModel: newAsProcModelModel(conn, c),
	}
}
