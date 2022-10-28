package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppApplyModel = (*customAsMarketAppApplyModel)(nil)

type (
	// AsMarketAppApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppApplyModel.
	AsMarketAppApplyModel interface {
		asMarketAppApplyModel
	}

	customAsMarketAppApplyModel struct {
		*defaultAsMarketAppApplyModel
	}
)

// NewAsMarketAppApplyModel returns a model for the database table.
func NewAsMarketAppApplyModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppApplyModel {
	return &customAsMarketAppApplyModel{
		defaultAsMarketAppApplyModel: newAsMarketAppApplyModel(conn, c),
	}
}
