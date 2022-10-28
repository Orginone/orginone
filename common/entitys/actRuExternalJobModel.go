package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuExternalJobModel = (*customActRuExternalJobModel)(nil)

type (
	// ActRuExternalJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuExternalJobModel.
	ActRuExternalJobModel interface {
		actRuExternalJobModel
	}

	customActRuExternalJobModel struct {
		*defaultActRuExternalJobModel
	}
)

// NewActRuExternalJobModel returns a model for the database table.
func NewActRuExternalJobModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuExternalJobModel {
	return &customActRuExternalJobModel{
		defaultActRuExternalJobModel: newActRuExternalJobModel(conn, c),
	}
}
