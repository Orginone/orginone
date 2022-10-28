package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormFinalModel = (*customAsFormFinalModel)(nil)

type (
	// AsFormFinalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormFinalModel.
	AsFormFinalModel interface {
		asFormFinalModel
	}

	customAsFormFinalModel struct {
		*defaultAsFormFinalModel
	}
)

// NewAsFormFinalModel returns a model for the database table.
func NewAsFormFinalModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormFinalModel {
	return &customAsFormFinalModel{
		defaultAsFormFinalModel: newAsFormFinalModel(conn, c),
	}
}
