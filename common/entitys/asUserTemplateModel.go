package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserTemplateModel = (*customAsUserTemplateModel)(nil)

type (
	// AsUserTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserTemplateModel.
	AsUserTemplateModel interface {
		asUserTemplateModel
	}

	customAsUserTemplateModel struct {
		*defaultAsUserTemplateModel
	}
)

// NewAsUserTemplateModel returns a model for the database table.
func NewAsUserTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserTemplateModel {
	return &customAsUserTemplateModel{
		defaultAsUserTemplateModel: newAsUserTemplateModel(conn, c),
	}
}
