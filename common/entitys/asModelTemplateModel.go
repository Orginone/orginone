package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsModelTemplateModel = (*customAsModelTemplateModel)(nil)

type (
	// AsModelTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsModelTemplateModel.
	AsModelTemplateModel interface {
		asModelTemplateModel
	}

	customAsModelTemplateModel struct {
		*defaultAsModelTemplateModel
	}
)

// NewAsModelTemplateModel returns a model for the database table.
func NewAsModelTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) AsModelTemplateModel {
	return &customAsModelTemplateModel{
		defaultAsModelTemplateModel: newAsModelTemplateModel(conn, c),
	}
}
