package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSysUserModel = (*customTSysUserModel)(nil)

type (
	// TSysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSysUserModel.
	TSysUserModel interface {
		tSysUserModel
	}

	customTSysUserModel struct {
		*defaultTSysUserModel
	}
)

// NewTSysUserModel returns a model for the database table.
func NewTSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) TSysUserModel {
	return &customTSysUserModel{
		defaultTSysUserModel: newTSysUserModel(conn, c),
	}
}
