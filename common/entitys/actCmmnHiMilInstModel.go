package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnHiMilInstModel = (*customActCmmnHiMilInstModel)(nil)

type (
	// ActCmmnHiMilInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnHiMilInstModel.
	ActCmmnHiMilInstModel interface {
		actCmmnHiMilInstModel
	}

	customActCmmnHiMilInstModel struct {
		*defaultActCmmnHiMilInstModel
	}
)

// NewActCmmnHiMilInstModel returns a model for the database table.
func NewActCmmnHiMilInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnHiMilInstModel {
	return &customActCmmnHiMilInstModel{
		defaultActCmmnHiMilInstModel: newActCmmnHiMilInstModel(conn, c),
	}
}
