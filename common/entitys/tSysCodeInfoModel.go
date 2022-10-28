package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSysCodeInfoModel = (*customTSysCodeInfoModel)(nil)

type (
	// TSysCodeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSysCodeInfoModel.
	TSysCodeInfoModel interface {
		tSysCodeInfoModel
	}

	customTSysCodeInfoModel struct {
		*defaultTSysCodeInfoModel
	}
)

// NewTSysCodeInfoModel returns a model for the database table.
func NewTSysCodeInfoModel(conn sqlx.SqlConn, c cache.CacheConf) TSysCodeInfoModel {
	return &customTSysCodeInfoModel{
		defaultTSysCodeInfoModel: newTSysCodeInfoModel(conn, c),
	}
}
