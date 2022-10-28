package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppPurchaseModel = (*customAsMarketAppPurchaseModel)(nil)

type (
	// AsMarketAppPurchaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppPurchaseModel.
	AsMarketAppPurchaseModel interface {
		asMarketAppPurchaseModel
	}

	customAsMarketAppPurchaseModel struct {
		*defaultAsMarketAppPurchaseModel
	}
)

// NewAsMarketAppPurchaseModel returns a model for the database table.
func NewAsMarketAppPurchaseModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppPurchaseModel {
	return &customAsMarketAppPurchaseModel{
		defaultAsMarketAppPurchaseModel: newAsMarketAppPurchaseModel(conn, c),
	}
}
