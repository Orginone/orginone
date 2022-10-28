package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsSmsMessageModel = (*customAsSmsMessageModel)(nil)

type (
	// AsSmsMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsSmsMessageModel.
	AsSmsMessageModel interface {
		asSmsMessageModel
	}

	customAsSmsMessageModel struct {
		*defaultAsSmsMessageModel
	}
)

// NewAsSmsMessageModel returns a model for the database table.
func NewAsSmsMessageModel(conn sqlx.SqlConn, c cache.CacheConf) AsSmsMessageModel {
	return &customAsSmsMessageModel{
		defaultAsSmsMessageModel: newAsSmsMessageModel(conn, c),
	}
}
