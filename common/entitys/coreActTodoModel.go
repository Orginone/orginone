package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreActTodoModel = (*customCoreActTodoModel)(nil)

type (
	// CoreActTodoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreActTodoModel.
	CoreActTodoModel interface {
		coreActTodoModel
	}

	customCoreActTodoModel struct {
		*defaultCoreActTodoModel
	}
)

// NewCoreActTodoModel returns a model for the database table.
func NewCoreActTodoModel(conn sqlx.SqlConn, c cache.CacheConf) CoreActTodoModel {
	return &customCoreActTodoModel{
		defaultCoreActTodoModel: newCoreActTodoModel(conn, c),
	}
}
