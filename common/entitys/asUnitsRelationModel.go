package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUnitsRelationModel = (*customAsUnitsRelationModel)(nil)

type (
	// AsUnitsRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUnitsRelationModel.
	AsUnitsRelationModel interface {
		asUnitsRelationModel
	}

	customAsUnitsRelationModel struct {
		*defaultAsUnitsRelationModel
	}
)

// NewAsUnitsRelationModel returns a model for the database table.
func NewAsUnitsRelationModel(conn sqlx.SqlConn, c cache.CacheConf) AsUnitsRelationModel {
	return &customAsUnitsRelationModel{
		defaultAsUnitsRelationModel: newAsUnitsRelationModel(conn, c),
	}
}
