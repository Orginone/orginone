package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsAppJobModel = (*customAsAppJobModel)(nil)

type (
	// AsAppJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsAppJobModel.
	AsAppJobModel interface {
		asAppJobModel
	}

	customAsAppJobModel struct {
		*defaultAsAppJobModel
	}
)

// NewAsAppJobModel returns a model for the database table.
func NewAsAppJobModel(conn sqlx.SqlConn, c cache.CacheConf) AsAppJobModel {
	return &customAsAppJobModel{
		defaultAsAppJobModel: newAsAppJobModel(conn, c),
	}
}
