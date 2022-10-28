package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuEventSubscrModel = (*customActRuEventSubscrModel)(nil)

type (
	// ActRuEventSubscrModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuEventSubscrModel.
	ActRuEventSubscrModel interface {
		actRuEventSubscrModel
	}

	customActRuEventSubscrModel struct {
		*defaultActRuEventSubscrModel
	}
)

// NewActRuEventSubscrModel returns a model for the database table.
func NewActRuEventSubscrModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuEventSubscrModel {
	return &customActRuEventSubscrModel{
		defaultActRuEventSubscrModel: newActRuEventSubscrModel(conn, c),
	}
}
