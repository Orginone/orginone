package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnHiPlanItemInstModel = (*customActCmmnHiPlanItemInstModel)(nil)

type (
	// ActCmmnHiPlanItemInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnHiPlanItemInstModel.
	ActCmmnHiPlanItemInstModel interface {
		actCmmnHiPlanItemInstModel
	}

	customActCmmnHiPlanItemInstModel struct {
		*defaultActCmmnHiPlanItemInstModel
	}
)

// NewActCmmnHiPlanItemInstModel returns a model for the database table.
func NewActCmmnHiPlanItemInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnHiPlanItemInstModel {
	return &customActCmmnHiPlanItemInstModel{
		defaultActCmmnHiPlanItemInstModel: newActCmmnHiPlanItemInstModel(conn, c),
	}
}
