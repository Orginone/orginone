package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDmnHiDecisionExecutionModel = (*customActDmnHiDecisionExecutionModel)(nil)

type (
	// ActDmnHiDecisionExecutionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDmnHiDecisionExecutionModel.
	ActDmnHiDecisionExecutionModel interface {
		actDmnHiDecisionExecutionModel
	}

	customActDmnHiDecisionExecutionModel struct {
		*defaultActDmnHiDecisionExecutionModel
	}
)

// NewActDmnHiDecisionExecutionModel returns a model for the database table.
func NewActDmnHiDecisionExecutionModel(conn sqlx.SqlConn, c cache.CacheConf) ActDmnHiDecisionExecutionModel {
	return &customActDmnHiDecisionExecutionModel{
		defaultActDmnHiDecisionExecutionModel: newActDmnHiDecisionExecutionModel(conn, c),
	}
}
