package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TFlowableFormModel = (*customTFlowableFormModel)(nil)

type (
	// TFlowableFormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFlowableFormModel.
	TFlowableFormModel interface {
		tFlowableFormModel
	}

	customTFlowableFormModel struct {
		*defaultTFlowableFormModel
	}
)

// NewTFlowableFormModel returns a model for the database table.
func NewTFlowableFormModel(conn sqlx.SqlConn, c cache.CacheConf) TFlowableFormModel {
	return &customTFlowableFormModel{
		defaultTFlowableFormModel: newTFlowableFormModel(conn, c),
	}
}
