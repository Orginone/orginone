package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CorePortalcontainModel = (*customCorePortalcontainModel)(nil)

type (
	// CorePortalcontainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCorePortalcontainModel.
	CorePortalcontainModel interface {
		corePortalcontainModel
	}

	customCorePortalcontainModel struct {
		*defaultCorePortalcontainModel
	}
)

// NewCorePortalcontainModel returns a model for the database table.
func NewCorePortalcontainModel(conn sqlx.SqlConn, c cache.CacheConf) CorePortalcontainModel {
	return &customCorePortalcontainModel{
		defaultCorePortalcontainModel: newCorePortalcontainModel(conn, c),
	}
}
