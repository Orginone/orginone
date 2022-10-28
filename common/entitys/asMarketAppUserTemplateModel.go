package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppUserTemplateModel = (*customAsMarketAppUserTemplateModel)(nil)

type (
	// AsMarketAppUserTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppUserTemplateModel.
	AsMarketAppUserTemplateModel interface {
		asMarketAppUserTemplateModel
	}

	customAsMarketAppUserTemplateModel struct {
		*defaultAsMarketAppUserTemplateModel
	}
)

// NewAsMarketAppUserTemplateModel returns a model for the database table.
func NewAsMarketAppUserTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppUserTemplateModel {
	return &customAsMarketAppUserTemplateModel{
		defaultAsMarketAppUserTemplateModel: newAsMarketAppUserTemplateModel(conn, c),
	}
}
