package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppComponentModel = (*customAsMarketAppComponentModel)(nil)

type (
	// AsMarketAppComponentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppComponentModel.
	AsMarketAppComponentModel interface {
		asMarketAppComponentModel
	}

	customAsMarketAppComponentModel struct {
		*defaultAsMarketAppComponentModel
	}
)

// NewAsMarketAppComponentModel returns a model for the database table.
func NewAsMarketAppComponentModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppComponentModel {
	return &customAsMarketAppComponentModel{
		defaultAsMarketAppComponentModel: newAsMarketAppComponentModel(conn, c),
	}
}
