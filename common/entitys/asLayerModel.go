package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsLayerModel = (*customAsLayerModel)(nil)

type (
	// AsLayerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsLayerModel.
	AsLayerModel interface {
		asLayerModel
	}

	customAsLayerModel struct {
		*defaultAsLayerModel
	}
)

// NewAsLayerModel returns a model for the database table.
func NewAsLayerModel(conn sqlx.SqlConn, c cache.CacheConf) AsLayerModel {
	return &customAsLayerModel{
		defaultAsLayerModel: newAsLayerModel(conn, c),
	}
}
