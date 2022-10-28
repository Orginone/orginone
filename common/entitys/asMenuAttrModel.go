package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMenuAttrModel = (*customAsMenuAttrModel)(nil)

type (
	// AsMenuAttrModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMenuAttrModel.
	AsMenuAttrModel interface {
		asMenuAttrModel
	}

	customAsMenuAttrModel struct {
		*defaultAsMenuAttrModel
	}
)

// NewAsMenuAttrModel returns a model for the database table.
func NewAsMenuAttrModel(conn sqlx.SqlConn, c cache.CacheConf) AsMenuAttrModel {
	return &customAsMenuAttrModel{
		defaultAsMenuAttrModel: newAsMenuAttrModel(conn, c),
	}
}
