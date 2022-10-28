package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsNbUnitModel = (*customAsNbUnitModel)(nil)

type (
	// AsNbUnitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsNbUnitModel.
	AsNbUnitModel interface {
		asNbUnitModel
	}

	customAsNbUnitModel struct {
		*defaultAsNbUnitModel
	}
)

// NewAsNbUnitModel returns a model for the database table.
func NewAsNbUnitModel(conn sqlx.SqlConn, c cache.CacheConf) AsNbUnitModel {
	return &customAsNbUnitModel{
		defaultAsNbUnitModel: newAsNbUnitModel(conn, c),
	}
}
