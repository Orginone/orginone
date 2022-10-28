package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnRuCaseInstModel = (*customActCmmnRuCaseInstModel)(nil)

type (
	// ActCmmnRuCaseInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnRuCaseInstModel.
	ActCmmnRuCaseInstModel interface {
		actCmmnRuCaseInstModel
	}

	customActCmmnRuCaseInstModel struct {
		*defaultActCmmnRuCaseInstModel
	}
)

// NewActCmmnRuCaseInstModel returns a model for the database table.
func NewActCmmnRuCaseInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnRuCaseInstModel {
	return &customActCmmnRuCaseInstModel{
		defaultActCmmnRuCaseInstModel: newActCmmnRuCaseInstModel(conn, c),
	}
}
