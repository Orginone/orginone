package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketRoleDistributionModel = (*customAsMarketRoleDistributionModel)(nil)

type (
	// AsMarketRoleDistributionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketRoleDistributionModel.
	AsMarketRoleDistributionModel interface {
		asMarketRoleDistributionModel
	}

	customAsMarketRoleDistributionModel struct {
		*defaultAsMarketRoleDistributionModel
	}
)

// NewAsMarketRoleDistributionModel returns a model for the database table.
func NewAsMarketRoleDistributionModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketRoleDistributionModel {
	return &customAsMarketRoleDistributionModel{
		defaultAsMarketRoleDistributionModel: newAsMarketRoleDistributionModel(conn, c),
	}
}
