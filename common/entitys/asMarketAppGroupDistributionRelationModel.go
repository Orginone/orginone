package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppGroupDistributionRelationModel = (*customAsMarketAppGroupDistributionRelationModel)(nil)

type (
	// AsMarketAppGroupDistributionRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppGroupDistributionRelationModel.
	AsMarketAppGroupDistributionRelationModel interface {
		asMarketAppGroupDistributionRelationModel
	}

	customAsMarketAppGroupDistributionRelationModel struct {
		*defaultAsMarketAppGroupDistributionRelationModel
	}
)

// NewAsMarketAppGroupDistributionRelationModel returns a model for the database table.
func NewAsMarketAppGroupDistributionRelationModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppGroupDistributionRelationModel {
	return &customAsMarketAppGroupDistributionRelationModel{
		defaultAsMarketAppGroupDistributionRelationModel: newAsMarketAppGroupDistributionRelationModel(conn, c),
	}
}
