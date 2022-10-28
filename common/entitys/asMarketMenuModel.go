package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketMenuModel = (*customAsMarketMenuModel)(nil)

type (
	// AsMarketMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketMenuModel.
	AsMarketMenuModel interface {
		asMarketMenuModel
	}

	customAsMarketMenuModel struct {
		*defaultAsMarketMenuModel
	}
)

// NewAsMarketMenuModel returns a model for the database table.
func NewAsMarketMenuModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketMenuModel {
	return &customAsMarketMenuModel{
		defaultAsMarketMenuModel: newAsMarketMenuModel(conn, c),
	}
}
