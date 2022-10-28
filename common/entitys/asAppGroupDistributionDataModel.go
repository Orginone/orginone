package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsAppGroupDistributionDataModel = (*customAsAppGroupDistributionDataModel)(nil)

type (
	// AsAppGroupDistributionDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsAppGroupDistributionDataModel.
	AsAppGroupDistributionDataModel interface {
		asAppGroupDistributionDataModel
	}

	customAsAppGroupDistributionDataModel struct {
		*defaultAsAppGroupDistributionDataModel
	}
)

// NewAsAppGroupDistributionDataModel returns a model for the database table.
func NewAsAppGroupDistributionDataModel(conn sqlx.SqlConn, c cache.CacheConf) AsAppGroupDistributionDataModel {
	return &customAsAppGroupDistributionDataModel{
		defaultAsAppGroupDistributionDataModel: newAsAppGroupDistributionDataModel(conn, c),
	}
}
