package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwEventDefinitionModel = (*customFlwEventDefinitionModel)(nil)

type (
	// FlwEventDefinitionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwEventDefinitionModel.
	FlwEventDefinitionModel interface {
		flwEventDefinitionModel
	}

	customFlwEventDefinitionModel struct {
		*defaultFlwEventDefinitionModel
	}
)

// NewFlwEventDefinitionModel returns a model for the database table.
func NewFlwEventDefinitionModel(conn sqlx.SqlConn, c cache.CacheConf) FlwEventDefinitionModel {
	return &customFlwEventDefinitionModel{
		defaultFlwEventDefinitionModel: newFlwEventDefinitionModel(conn, c),
	}
}
