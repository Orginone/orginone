package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsPropertiesDistributionModel = (*customAsPropertiesDistributionModel)(nil)

type (
	// AsPropertiesDistributionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsPropertiesDistributionModel.
	AsPropertiesDistributionModel interface {
		asPropertiesDistributionModel
	}

	customAsPropertiesDistributionModel struct {
		*defaultAsPropertiesDistributionModel
	}
)

// NewAsPropertiesDistributionModel returns a model for the database table.
func NewAsPropertiesDistributionModel(conn sqlx.SqlConn, c cache.CacheConf) AsPropertiesDistributionModel {
	return &customAsPropertiesDistributionModel{
		defaultAsPropertiesDistributionModel: newAsPropertiesDistributionModel(conn, c),
	}
}
