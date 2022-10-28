package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsApplicationModel = (*customAsApplicationModel)(nil)

type (
	// AsApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsApplicationModel.
	AsApplicationModel interface {
		asApplicationModel
	}

	customAsApplicationModel struct {
		*defaultAsApplicationModel
	}
)

// NewAsApplicationModel returns a model for the database table.
func NewAsApplicationModel(conn sqlx.SqlConn, c cache.CacheConf) AsApplicationModel {
	return &customAsApplicationModel{
		defaultAsApplicationModel: newAsApplicationModel(conn, c),
	}
}
