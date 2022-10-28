package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMessageReadRecordModel = (*customAsMessageReadRecordModel)(nil)

type (
	// AsMessageReadRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMessageReadRecordModel.
	AsMessageReadRecordModel interface {
		asMessageReadRecordModel
	}

	customAsMessageReadRecordModel struct {
		*defaultAsMessageReadRecordModel
	}
)

// NewAsMessageReadRecordModel returns a model for the database table.
func NewAsMessageReadRecordModel(conn sqlx.SqlConn, c cache.CacheConf) AsMessageReadRecordModel {
	return &customAsMessageReadRecordModel{
		defaultAsMessageReadRecordModel: newAsMessageReadRecordModel(conn, c),
	}
}
