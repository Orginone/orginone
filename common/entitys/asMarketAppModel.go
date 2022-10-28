package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppModel = (*customAsMarketAppModel)(nil)

type (
	// AsMarketAppModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppModel.
	AsMarketAppModel interface {
		asMarketAppModel
	}

	customAsMarketAppModel struct {
		*defaultAsMarketAppModel
	}
)

// NewAsMarketAppModel returns a model for the database table.
func NewAsMarketAppModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppModel {
	return &customAsMarketAppModel{
		defaultAsMarketAppModel: newAsMarketAppModel(conn, c),
	}
}
