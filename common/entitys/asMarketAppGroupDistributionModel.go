package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppGroupDistributionModel = (*customAsMarketAppGroupDistributionModel)(nil)

type (
	// AsMarketAppGroupDistributionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppGroupDistributionModel.
	AsMarketAppGroupDistributionModel interface {
		asMarketAppGroupDistributionModel
	}

	customAsMarketAppGroupDistributionModel struct {
		*defaultAsMarketAppGroupDistributionModel
	}
)

// NewAsMarketAppGroupDistributionModel returns a model for the database table.
func NewAsMarketAppGroupDistributionModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppGroupDistributionModel {
	return &customAsMarketAppGroupDistributionModel{
		defaultAsMarketAppGroupDistributionModel: newAsMarketAppGroupDistributionModel(conn, c),
	}
}
