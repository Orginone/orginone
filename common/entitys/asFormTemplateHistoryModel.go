package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormTemplateHistoryModel = (*customAsFormTemplateHistoryModel)(nil)

type (
	// AsFormTemplateHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormTemplateHistoryModel.
	AsFormTemplateHistoryModel interface {
		asFormTemplateHistoryModel
	}

	customAsFormTemplateHistoryModel struct {
		*defaultAsFormTemplateHistoryModel
	}
)

// NewAsFormTemplateHistoryModel returns a model for the database table.
func NewAsFormTemplateHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormTemplateHistoryModel {
	return &customAsFormTemplateHistoryModel{
		defaultAsFormTemplateHistoryModel: newAsFormTemplateHistoryModel(conn, c),
	}
}
