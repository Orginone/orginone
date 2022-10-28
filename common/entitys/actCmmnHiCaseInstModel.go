package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnHiCaseInstModel = (*customActCmmnHiCaseInstModel)(nil)

type (
	// ActCmmnHiCaseInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnHiCaseInstModel.
	ActCmmnHiCaseInstModel interface {
		actCmmnHiCaseInstModel
	}

	customActCmmnHiCaseInstModel struct {
		*defaultActCmmnHiCaseInstModel
	}
)

// NewActCmmnHiCaseInstModel returns a model for the database table.
func NewActCmmnHiCaseInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnHiCaseInstModel {
	return &customActCmmnHiCaseInstModel{
		defaultActCmmnHiCaseInstModel: newActCmmnHiCaseInstModel(conn, c),
	}
}
