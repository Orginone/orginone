package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketUsedAppModel = (*customAsMarketUsedAppModel)(nil)

type (
	// AsMarketUsedAppModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketUsedAppModel.
	AsMarketUsedAppModel interface {
		asMarketUsedAppModel
	}

	customAsMarketUsedAppModel struct {
		*defaultAsMarketUsedAppModel
	}
)

// NewAsMarketUsedAppModel returns a model for the database table.
func NewAsMarketUsedAppModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketUsedAppModel {
	return &customAsMarketUsedAppModel{
		defaultAsMarketUsedAppModel: newAsMarketUsedAppModel(conn, c),
	}
}
