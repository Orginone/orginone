package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppAlertModel = (*customAsMarketAppAlertModel)(nil)

type (
	// AsMarketAppAlertModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppAlertModel.
	AsMarketAppAlertModel interface {
		asMarketAppAlertModel
	}

	customAsMarketAppAlertModel struct {
		*defaultAsMarketAppAlertModel
	}
)

// NewAsMarketAppAlertModel returns a model for the database table.
func NewAsMarketAppAlertModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppAlertModel {
	return &customAsMarketAppAlertModel{
		defaultAsMarketAppAlertModel: newAsMarketAppAlertModel(conn, c),
	}
}
