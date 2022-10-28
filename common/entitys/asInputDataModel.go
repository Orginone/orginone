package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsInputDataModel = (*customAsInputDataModel)(nil)

type (
	// AsInputDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsInputDataModel.
	AsInputDataModel interface {
		asInputDataModel
	}

	customAsInputDataModel struct {
		*defaultAsInputDataModel
	}
)

// NewAsInputDataModel returns a model for the database table.
func NewAsInputDataModel(conn sqlx.SqlConn, c cache.CacheConf) AsInputDataModel {
	return &customAsInputDataModel{
		defaultAsInputDataModel: newAsInputDataModel(conn, c),
	}
}
