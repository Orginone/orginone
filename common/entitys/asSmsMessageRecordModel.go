package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsSmsMessageRecordModel = (*customAsSmsMessageRecordModel)(nil)

type (
	// AsSmsMessageRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsSmsMessageRecordModel.
	AsSmsMessageRecordModel interface {
		asSmsMessageRecordModel
	}

	customAsSmsMessageRecordModel struct {
		*defaultAsSmsMessageRecordModel
	}
)

// NewAsSmsMessageRecordModel returns a model for the database table.
func NewAsSmsMessageRecordModel(conn sqlx.SqlConn, c cache.CacheConf) AsSmsMessageRecordModel {
	return &customAsSmsMessageRecordModel{
		defaultAsSmsMessageRecordModel: newAsSmsMessageRecordModel(conn, c),
	}
}
