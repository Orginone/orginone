package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTaskModel = (*customAsTaskModel)(nil)

type (
	// AsTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTaskModel.
	AsTaskModel interface {
		asTaskModel
	}

	customAsTaskModel struct {
		*defaultAsTaskModel
	}
)

// NewAsTaskModel returns a model for the database table.
func NewAsTaskModel(conn sqlx.SqlConn, c cache.CacheConf) AsTaskModel {
	return &customAsTaskModel{
		defaultAsTaskModel: newAsTaskModel(conn, c),
	}
}
