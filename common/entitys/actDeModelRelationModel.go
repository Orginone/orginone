package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDeModelRelationModel = (*customActDeModelRelationModel)(nil)

type (
	// ActDeModelRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDeModelRelationModel.
	ActDeModelRelationModel interface {
		actDeModelRelationModel
	}

	customActDeModelRelationModel struct {
		*defaultActDeModelRelationModel
	}
)

// NewActDeModelRelationModel returns a model for the database table.
func NewActDeModelRelationModel(conn sqlx.SqlConn, c cache.CacheConf) ActDeModelRelationModel {
	return &customActDeModelRelationModel{
		defaultActDeModelRelationModel: newActDeModelRelationModel(conn, c),
	}
}
