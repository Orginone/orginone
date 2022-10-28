package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserFormModel = (*customAsUserFormModel)(nil)

type (
	// AsUserFormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserFormModel.
	AsUserFormModel interface {
		asUserFormModel
	}

	customAsUserFormModel struct {
		*defaultAsUserFormModel
	}
)

// NewAsUserFormModel returns a model for the database table.
func NewAsUserFormModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserFormModel {
	return &customAsUserFormModel{
		defaultAsUserFormModel: newAsUserFormModel(conn, c),
	}
}
