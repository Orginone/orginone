package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MarketPortalModel = (*customMarketPortalModel)(nil)

type (
	// MarketPortalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMarketPortalModel.
	MarketPortalModel interface {
		marketPortalModel
	}

	customMarketPortalModel struct {
		*defaultMarketPortalModel
	}
)

// NewMarketPortalModel returns a model for the database table.
func NewMarketPortalModel(conn sqlx.SqlConn, c cache.CacheConf) MarketPortalModel {
	return &customMarketPortalModel{
		defaultMarketPortalModel: newMarketPortalModel(conn, c),
	}
}
