package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TempModel = (*customTempModel)(nil)

type (
	// TempModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTempModel.
	TempModel interface {
		tempModel
	}

	customTempModel struct {
		*defaultTempModel
	}
)

// NewTempModel returns a model for the database table.
func NewTempModel(conn sqlx.SqlConn, c cache.CacheConf) TempModel {
	return &customTempModel{
		defaultTempModel: newTempModel(conn, c),
	}
}
