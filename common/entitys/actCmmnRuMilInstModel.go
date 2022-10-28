package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnRuMilInstModel = (*customActCmmnRuMilInstModel)(nil)

type (
	// ActCmmnRuMilInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnRuMilInstModel.
	ActCmmnRuMilInstModel interface {
		actCmmnRuMilInstModel
	}

	customActCmmnRuMilInstModel struct {
		*defaultActCmmnRuMilInstModel
	}
)

// NewActCmmnRuMilInstModel returns a model for the database table.
func NewActCmmnRuMilInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnRuMilInstModel {
	return &customActCmmnRuMilInstModel{
		defaultActCmmnRuMilInstModel: newActCmmnRuMilInstModel(conn, c),
	}
}
