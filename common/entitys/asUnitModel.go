package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUnitModel = (*customAsUnitModel)(nil)

type (
	// AsUnitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUnitModel.
	AsUnitModel interface {
		asUnitModel
	}

	customAsUnitModel struct {
		*defaultAsUnitModel
	}
)

// NewAsUnitModel returns a model for the database table.
func NewAsUnitModel(conn sqlx.SqlConn, c cache.CacheConf) AsUnitModel {
	return &customAsUnitModel{
		defaultAsUnitModel: newAsUnitModel(conn, c),
	}
}
