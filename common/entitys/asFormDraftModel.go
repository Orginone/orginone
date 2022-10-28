package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormDraftModel = (*customAsFormDraftModel)(nil)

type (
	// AsFormDraftModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormDraftModel.
	AsFormDraftModel interface {
		asFormDraftModel
	}

	customAsFormDraftModel struct {
		*defaultAsFormDraftModel
	}
)

// NewAsFormDraftModel returns a model for the database table.
func NewAsFormDraftModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormDraftModel {
	return &customAsFormDraftModel{
		defaultAsFormDraftModel: newAsFormDraftModel(conn, c),
	}
}
