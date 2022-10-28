package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreZcCardstyleModel = (*customCoreZcCardstyleModel)(nil)

type (
	// CoreZcCardstyleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreZcCardstyleModel.
	CoreZcCardstyleModel interface {
		coreZcCardstyleModel
	}

	customCoreZcCardstyleModel struct {
		*defaultCoreZcCardstyleModel
	}
)

// NewCoreZcCardstyleModel returns a model for the database table.
func NewCoreZcCardstyleModel(conn sqlx.SqlConn, c cache.CacheConf) CoreZcCardstyleModel {
	return &customCoreZcCardstyleModel{
		defaultCoreZcCardstyleModel: newCoreZcCardstyleModel(conn, c),
	}
}
