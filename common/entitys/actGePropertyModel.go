package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActGePropertyModel = (*customActGePropertyModel)(nil)

type (
	// ActGePropertyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActGePropertyModel.
	ActGePropertyModel interface {
		actGePropertyModel
	}

	customActGePropertyModel struct {
		*defaultActGePropertyModel
	}
)

// NewActGePropertyModel returns a model for the database table.
func NewActGePropertyModel(conn sqlx.SqlConn, c cache.CacheConf) ActGePropertyModel {
	return &customActGePropertyModel{
		defaultActGePropertyModel: newActGePropertyModel(conn, c),
	}
}
