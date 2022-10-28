package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreFieldModel = (*customCoreFieldModel)(nil)

type (
	// CoreFieldModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreFieldModel.
	CoreFieldModel interface {
		coreFieldModel
	}

	customCoreFieldModel struct {
		*defaultCoreFieldModel
	}
)

// NewCoreFieldModel returns a model for the database table.
func NewCoreFieldModel(conn sqlx.SqlConn, c cache.CacheConf) CoreFieldModel {
	return &customCoreFieldModel{
		defaultCoreFieldModel: newCoreFieldModel(conn, c),
	}
}
