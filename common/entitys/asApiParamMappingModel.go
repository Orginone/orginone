package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsApiParamMappingModel = (*customAsApiParamMappingModel)(nil)

type (
	// AsApiParamMappingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsApiParamMappingModel.
	AsApiParamMappingModel interface {
		asApiParamMappingModel
	}

	customAsApiParamMappingModel struct {
		*defaultAsApiParamMappingModel
	}
)

// NewAsApiParamMappingModel returns a model for the database table.
func NewAsApiParamMappingModel(conn sqlx.SqlConn, c cache.CacheConf) AsApiParamMappingModel {
	return &customAsApiParamMappingModel{
		defaultAsApiParamMappingModel: newAsApiParamMappingModel(conn, c),
	}
}
