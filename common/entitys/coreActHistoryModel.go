package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreActHistoryModel = (*customCoreActHistoryModel)(nil)

type (
	// CoreActHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreActHistoryModel.
	CoreActHistoryModel interface {
		coreActHistoryModel
	}

	customCoreActHistoryModel struct {
		*defaultCoreActHistoryModel
	}
)

// NewCoreActHistoryModel returns a model for the database table.
func NewCoreActHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) CoreActHistoryModel {
	return &customCoreActHistoryModel{
		defaultCoreActHistoryModel: newCoreActHistoryModel(conn, c),
	}
}
