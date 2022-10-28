package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsWorkingDataModel = (*customAsWorkingDataModel)(nil)

type (
	// AsWorkingDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsWorkingDataModel.
	AsWorkingDataModel interface {
		asWorkingDataModel
	}

	customAsWorkingDataModel struct {
		*defaultAsWorkingDataModel
	}
)

// NewAsWorkingDataModel returns a model for the database table.
func NewAsWorkingDataModel(conn sqlx.SqlConn, c cache.CacheConf) AsWorkingDataModel {
	return &customAsWorkingDataModel{
		defaultAsWorkingDataModel: newAsWorkingDataModel(conn, c),
	}
}
