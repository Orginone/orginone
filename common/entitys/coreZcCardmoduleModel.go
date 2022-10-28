package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreZcCardmoduleModel = (*customCoreZcCardmoduleModel)(nil)

type (
	// CoreZcCardmoduleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreZcCardmoduleModel.
	CoreZcCardmoduleModel interface {
		coreZcCardmoduleModel
	}

	customCoreZcCardmoduleModel struct {
		*defaultCoreZcCardmoduleModel
	}
)

// NewCoreZcCardmoduleModel returns a model for the database table.
func NewCoreZcCardmoduleModel(conn sqlx.SqlConn, c cache.CacheConf) CoreZcCardmoduleModel {
	return &customCoreZcCardmoduleModel{
		defaultCoreZcCardmoduleModel: newCoreZcCardmoduleModel(conn, c),
	}
}
