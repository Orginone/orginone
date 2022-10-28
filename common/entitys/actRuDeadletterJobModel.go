package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuDeadletterJobModel = (*customActRuDeadletterJobModel)(nil)

type (
	// ActRuDeadletterJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuDeadletterJobModel.
	ActRuDeadletterJobModel interface {
		actRuDeadletterJobModel
	}

	customActRuDeadletterJobModel struct {
		*defaultActRuDeadletterJobModel
	}
)

// NewActRuDeadletterJobModel returns a model for the database table.
func NewActRuDeadletterJobModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuDeadletterJobModel {
	return &customActRuDeadletterJobModel{
		defaultActRuDeadletterJobModel: newActRuDeadletterJobModel(conn, c),
	}
}
