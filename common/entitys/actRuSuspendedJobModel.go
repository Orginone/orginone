package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuSuspendedJobModel = (*customActRuSuspendedJobModel)(nil)

type (
	// ActRuSuspendedJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuSuspendedJobModel.
	ActRuSuspendedJobModel interface {
		actRuSuspendedJobModel
	}

	customActRuSuspendedJobModel struct {
		*defaultActRuSuspendedJobModel
	}
)

// NewActRuSuspendedJobModel returns a model for the database table.
func NewActRuSuspendedJobModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuSuspendedJobModel {
	return &customActRuSuspendedJobModel{
		defaultActRuSuspendedJobModel: newActRuSuspendedJobModel(conn, c),
	}
}
