package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwRuBatchModel = (*customFlwRuBatchModel)(nil)

type (
	// FlwRuBatchModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwRuBatchModel.
	FlwRuBatchModel interface {
		flwRuBatchModel
	}

	customFlwRuBatchModel struct {
		*defaultFlwRuBatchModel
	}
)

// NewFlwRuBatchModel returns a model for the database table.
func NewFlwRuBatchModel(conn sqlx.SqlConn, c cache.CacheConf) FlwRuBatchModel {
	return &customFlwRuBatchModel{
		defaultFlwRuBatchModel: newFlwRuBatchModel(conn, c),
	}
}
