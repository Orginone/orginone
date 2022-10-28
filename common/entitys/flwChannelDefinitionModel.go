package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwChannelDefinitionModel = (*customFlwChannelDefinitionModel)(nil)

type (
	// FlwChannelDefinitionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwChannelDefinitionModel.
	FlwChannelDefinitionModel interface {
		flwChannelDefinitionModel
	}

	customFlwChannelDefinitionModel struct {
		*defaultFlwChannelDefinitionModel
	}
)

// NewFlwChannelDefinitionModel returns a model for the database table.
func NewFlwChannelDefinitionModel(conn sqlx.SqlConn, c cache.CacheConf) FlwChannelDefinitionModel {
	return &customFlwChannelDefinitionModel{
		defaultFlwChannelDefinitionModel: newFlwChannelDefinitionModel(conn, c),
	}
}
