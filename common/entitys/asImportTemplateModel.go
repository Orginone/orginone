package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsImportTemplateModel = (*customAsImportTemplateModel)(nil)

type (
	// AsImportTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsImportTemplateModel.
	AsImportTemplateModel interface {
		asImportTemplateModel
	}

	customAsImportTemplateModel struct {
		*defaultAsImportTemplateModel
	}
)

// NewAsImportTemplateModel returns a model for the database table.
func NewAsImportTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) AsImportTemplateModel {
	return &customAsImportTemplateModel{
		defaultAsImportTemplateModel: newAsImportTemplateModel(conn, c),
	}
}
