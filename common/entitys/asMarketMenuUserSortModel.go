package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketMenuUserSortModel = (*customAsMarketMenuUserSortModel)(nil)

type (
	// AsMarketMenuUserSortModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketMenuUserSortModel.
	AsMarketMenuUserSortModel interface {
		asMarketMenuUserSortModel
	}

	customAsMarketMenuUserSortModel struct {
		*defaultAsMarketMenuUserSortModel
	}
)

// NewAsMarketMenuUserSortModel returns a model for the database table.
func NewAsMarketMenuUserSortModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketMenuUserSortModel {
	return &customAsMarketMenuUserSortModel{
		defaultAsMarketMenuUserSortModel: newAsMarketMenuUserSortModel(conn, c),
	}
}
