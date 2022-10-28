package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuTimerJobModel = (*customActRuTimerJobModel)(nil)

type (
	// ActRuTimerJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuTimerJobModel.
	ActRuTimerJobModel interface {
		actRuTimerJobModel
	}

	customActRuTimerJobModel struct {
		*defaultActRuTimerJobModel
	}
)

// NewActRuTimerJobModel returns a model for the database table.
func NewActRuTimerJobModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuTimerJobModel {
	return &customActRuTimerJobModel{
		defaultActRuTimerJobModel: newActRuTimerJobModel(conn, c),
	}
}
