package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwRuBatchPartModel = (*customFlwRuBatchPartModel)(nil)

type (
	// FlwRuBatchPartModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwRuBatchPartModel.
	FlwRuBatchPartModel interface {
		flwRuBatchPartModel
	}

	customFlwRuBatchPartModel struct {
		*defaultFlwRuBatchPartModel
	}
)

// NewFlwRuBatchPartModel returns a model for the database table.
func NewFlwRuBatchPartModel(conn sqlx.SqlConn, c cache.CacheConf) FlwRuBatchPartModel {
	return &customFlwRuBatchPartModel{
		defaultFlwRuBatchPartModel: newFlwRuBatchPartModel(conn, c),
	}
}
