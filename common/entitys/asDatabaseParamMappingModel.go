package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsDatabaseParamMappingModel = (*customAsDatabaseParamMappingModel)(nil)

type (
	// AsDatabaseParamMappingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsDatabaseParamMappingModel.
	AsDatabaseParamMappingModel interface {
		asDatabaseParamMappingModel
	}

	customAsDatabaseParamMappingModel struct {
		*defaultAsDatabaseParamMappingModel
	}
)

// NewAsDatabaseParamMappingModel returns a model for the database table.
func NewAsDatabaseParamMappingModel(conn sqlx.SqlConn, c cache.CacheConf) AsDatabaseParamMappingModel {
	return &customAsDatabaseParamMappingModel{
		defaultAsDatabaseParamMappingModel: newAsDatabaseParamMappingModel(conn, c),
	}
}
