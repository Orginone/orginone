package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormTemplateModel = (*customAsFormTemplateModel)(nil)

type (
	// AsFormTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormTemplateModel.
	AsFormTemplateModel interface {
		asFormTemplateModel
	}

	customAsFormTemplateModel struct {
		*defaultAsFormTemplateModel
	}
)

// NewAsFormTemplateModel returns a model for the database table.
func NewAsFormTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormTemplateModel {
	return &customAsFormTemplateModel{
		defaultAsFormTemplateModel: newAsFormTemplateModel(conn, c),
	}
}
