package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserBindModel = (*customAsUserBindModel)(nil)

type (
	// AsUserBindModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserBindModel.
	AsUserBindModel interface {
		asUserBindModel
	}

	customAsUserBindModel struct {
		*defaultAsUserBindModel
	}
)

// NewAsUserBindModel returns a model for the database table.
func NewAsUserBindModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserBindModel {
	return &customAsUserBindModel{
		defaultAsUserBindModel: newAsUserBindModel(conn, c),
	}
}
