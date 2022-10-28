package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserLoginTimeModel = (*customAsUserLoginTimeModel)(nil)

type (
	// AsUserLoginTimeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserLoginTimeModel.
	AsUserLoginTimeModel interface {
		asUserLoginTimeModel
	}

	customAsUserLoginTimeModel struct {
		*defaultAsUserLoginTimeModel
	}
)

// NewAsUserLoginTimeModel returns a model for the database table.
func NewAsUserLoginTimeModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserLoginTimeModel {
	return &customAsUserLoginTimeModel{
		defaultAsUserLoginTimeModel: newAsUserLoginTimeModel(conn, c),
	}
}
