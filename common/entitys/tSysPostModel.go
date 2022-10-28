package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSysPostModel = (*customTSysPostModel)(nil)

type (
	// TSysPostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSysPostModel.
	TSysPostModel interface {
		tSysPostModel
	}

	customTSysPostModel struct {
		*defaultTSysPostModel
	}
)

// NewTSysPostModel returns a model for the database table.
func NewTSysPostModel(conn sqlx.SqlConn, c cache.CacheConf) TSysPostModel {
	return &customTSysPostModel{
		defaultTSysPostModel: newTSysPostModel(conn, c),
	}
}
