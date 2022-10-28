package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnRuPlanItemInstModel = (*customActCmmnRuPlanItemInstModel)(nil)

type (
	// ActCmmnRuPlanItemInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnRuPlanItemInstModel.
	ActCmmnRuPlanItemInstModel interface {
		actCmmnRuPlanItemInstModel
	}

	customActCmmnRuPlanItemInstModel struct {
		*defaultActCmmnRuPlanItemInstModel
	}
)

// NewActCmmnRuPlanItemInstModel returns a model for the database table.
func NewActCmmnRuPlanItemInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnRuPlanItemInstModel {
	return &customActCmmnRuPlanItemInstModel{
		defaultActCmmnRuPlanItemInstModel: newActCmmnRuPlanItemInstModel(conn, c),
	}
}
