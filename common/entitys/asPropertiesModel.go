package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsPropertiesModel = (*customAsPropertiesModel)(nil)

type (
	// AsPropertiesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsPropertiesModel.
	AsPropertiesModel interface {
		asPropertiesModel
	}

	customAsPropertiesModel struct {
		*defaultAsPropertiesModel
	}
)

// NewAsPropertiesModel returns a model for the database table.
func NewAsPropertiesModel(conn sqlx.SqlConn, c cache.CacheConf) AsPropertiesModel {
	return &customAsPropertiesModel{
		defaultAsPropertiesModel: newAsPropertiesModel(conn, c),
	}
}
