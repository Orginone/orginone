package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormModelModel = (*customAsFormModelModel)(nil)

type (
	// AsFormModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormModelModel.
	AsFormModelModel interface {
		asFormModelModel
	}

	customAsFormModelModel struct {
		*defaultAsFormModelModel
	}
)

// NewAsFormModelModel returns a model for the database table.
func NewAsFormModelModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormModelModel {
	return &customAsFormModelModel{
		defaultAsFormModelModel: newAsFormModelModel(conn, c),
	}
}
