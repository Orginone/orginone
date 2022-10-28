package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActFoFormDefinitionModel = (*customActFoFormDefinitionModel)(nil)

type (
	// ActFoFormDefinitionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActFoFormDefinitionModel.
	ActFoFormDefinitionModel interface {
		actFoFormDefinitionModel
	}

	customActFoFormDefinitionModel struct {
		*defaultActFoFormDefinitionModel
	}
)

// NewActFoFormDefinitionModel returns a model for the database table.
func NewActFoFormDefinitionModel(conn sqlx.SqlConn, c cache.CacheConf) ActFoFormDefinitionModel {
	return &customActFoFormDefinitionModel{
		defaultActFoFormDefinitionModel: newActFoFormDefinitionModel(conn, c),
	}
}
