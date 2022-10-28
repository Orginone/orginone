package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnRuSentryPartInstModel = (*customActCmmnRuSentryPartInstModel)(nil)

type (
	// ActCmmnRuSentryPartInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnRuSentryPartInstModel.
	ActCmmnRuSentryPartInstModel interface {
		actCmmnRuSentryPartInstModel
	}

	customActCmmnRuSentryPartInstModel struct {
		*defaultActCmmnRuSentryPartInstModel
	}
)

// NewActCmmnRuSentryPartInstModel returns a model for the database table.
func NewActCmmnRuSentryPartInstModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnRuSentryPartInstModel {
	return &customActCmmnRuSentryPartInstModel{
		defaultActCmmnRuSentryPartInstModel: newActCmmnRuSentryPartInstModel(conn, c),
	}
}
