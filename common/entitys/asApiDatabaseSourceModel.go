package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsApiDatabaseSourceModel = (*customAsApiDatabaseSourceModel)(nil)

type (
	// AsApiDatabaseSourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsApiDatabaseSourceModel.
	AsApiDatabaseSourceModel interface {
		asApiDatabaseSourceModel
	}

	customAsApiDatabaseSourceModel struct {
		*defaultAsApiDatabaseSourceModel
	}
)

// NewAsApiDatabaseSourceModel returns a model for the database table.
func NewAsApiDatabaseSourceModel(conn sqlx.SqlConn, c cache.CacheConf) AsApiDatabaseSourceModel {
	return &customAsApiDatabaseSourceModel{
		defaultAsApiDatabaseSourceModel: newAsApiDatabaseSourceModel(conn, c),
	}
}
