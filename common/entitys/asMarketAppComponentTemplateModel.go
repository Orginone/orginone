package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppComponentTemplateModel = (*customAsMarketAppComponentTemplateModel)(nil)

type (
	// AsMarketAppComponentTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppComponentTemplateModel.
	AsMarketAppComponentTemplateModel interface {
		asMarketAppComponentTemplateModel
	}

	customAsMarketAppComponentTemplateModel struct {
		*defaultAsMarketAppComponentTemplateModel
	}
)

// NewAsMarketAppComponentTemplateModel returns a model for the database table.
func NewAsMarketAppComponentTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppComponentTemplateModel {
	return &customAsMarketAppComponentTemplateModel{
		defaultAsMarketAppComponentTemplateModel: newAsMarketAppComponentTemplateModel(conn, c),
	}
}
