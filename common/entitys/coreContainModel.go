package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreContainModel = (*customCoreContainModel)(nil)

type (
	// CoreContainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreContainModel.
	CoreContainModel interface {
		coreContainModel
	}

	customCoreContainModel struct {
		*defaultCoreContainModel
	}
)

// NewCoreContainModel returns a model for the database table.
func NewCoreContainModel(conn sqlx.SqlConn, c cache.CacheConf) CoreContainModel {
	return &customCoreContainModel{
		defaultCoreContainModel: newCoreContainModel(conn, c),
	}
}
