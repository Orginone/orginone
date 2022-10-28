package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdPrivMappingModel = (*customActIdPrivMappingModel)(nil)

type (
	// ActIdPrivMappingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdPrivMappingModel.
	ActIdPrivMappingModel interface {
		actIdPrivMappingModel
	}

	customActIdPrivMappingModel struct {
		*defaultActIdPrivMappingModel
	}
)

// NewActIdPrivMappingModel returns a model for the database table.
func NewActIdPrivMappingModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdPrivMappingModel {
	return &customActIdPrivMappingModel{
		defaultActIdPrivMappingModel: newActIdPrivMappingModel(conn, c),
	}
}
