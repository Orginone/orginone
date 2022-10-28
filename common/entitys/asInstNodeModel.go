package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsInstNodeModel = (*customAsInstNodeModel)(nil)

type (
	// AsInstNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsInstNodeModel.
	AsInstNodeModel interface {
		asInstNodeModel
	}

	customAsInstNodeModel struct {
		*defaultAsInstNodeModel
	}
)

// NewAsInstNodeModel returns a model for the database table.
func NewAsInstNodeModel(conn sqlx.SqlConn, c cache.CacheConf) AsInstNodeModel {
	return &customAsInstNodeModel{
		defaultAsInstNodeModel: newAsInstNodeModel(conn, c),
	}
}
