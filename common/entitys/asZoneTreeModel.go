package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsZoneTreeModel = (*customAsZoneTreeModel)(nil)

type (
	// AsZoneTreeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsZoneTreeModel.
	AsZoneTreeModel interface {
		asZoneTreeModel
	}

	customAsZoneTreeModel struct {
		*defaultAsZoneTreeModel
	}
)

// NewAsZoneTreeModel returns a model for the database table.
func NewAsZoneTreeModel(conn sqlx.SqlConn, c cache.CacheConf) AsZoneTreeModel {
	return &customAsZoneTreeModel{
		defaultAsZoneTreeModel: newAsZoneTreeModel(conn, c),
	}
}
