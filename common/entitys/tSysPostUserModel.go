package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSysPostUserModel = (*customTSysPostUserModel)(nil)

type (
	// TSysPostUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSysPostUserModel.
	TSysPostUserModel interface {
		tSysPostUserModel
	}

	customTSysPostUserModel struct {
		*defaultTSysPostUserModel
	}
)

// NewTSysPostUserModel returns a model for the database table.
func NewTSysPostUserModel(conn sqlx.SqlConn, c cache.CacheConf) TSysPostUserModel {
	return &customTSysPostUserModel{
		defaultTSysPostUserModel: newTSysPostUserModel(conn, c),
	}
}
