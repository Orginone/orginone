package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsApiDataSourceModel = (*customAsApiDataSourceModel)(nil)

type (
	// AsApiDataSourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsApiDataSourceModel.
	AsApiDataSourceModel interface {
		asApiDataSourceModel
	}

	customAsApiDataSourceModel struct {
		*defaultAsApiDataSourceModel
	}
)

// NewAsApiDataSourceModel returns a model for the database table.
func NewAsApiDataSourceModel(conn sqlx.SqlConn, c cache.CacheConf) AsApiDataSourceModel {
	return &customAsApiDataSourceModel{
		defaultAsApiDataSourceModel: newAsApiDataSourceModel(conn, c),
	}
}
