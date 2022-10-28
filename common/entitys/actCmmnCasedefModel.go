package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnCasedefModel = (*customActCmmnCasedefModel)(nil)

type (
	// ActCmmnCasedefModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnCasedefModel.
	ActCmmnCasedefModel interface {
		actCmmnCasedefModel
	}

	customActCmmnCasedefModel struct {
		*defaultActCmmnCasedefModel
	}
)

// NewActCmmnCasedefModel returns a model for the database table.
func NewActCmmnCasedefModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnCasedefModel {
	return &customActCmmnCasedefModel{
		defaultActCmmnCasedefModel: newActCmmnCasedefModel(conn, c),
	}
}
