package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CorePortalModel = (*customCorePortalModel)(nil)

type (
	// CorePortalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCorePortalModel.
	CorePortalModel interface {
		corePortalModel
	}

	customCorePortalModel struct {
		*defaultCorePortalModel
	}
)

// NewCorePortalModel returns a model for the database table.
func NewCorePortalModel(conn sqlx.SqlConn, c cache.CacheConf) CorePortalModel {
	return &customCorePortalModel{
		defaultCorePortalModel: newCorePortalModel(conn, c),
	}
}
