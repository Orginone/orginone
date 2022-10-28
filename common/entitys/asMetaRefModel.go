package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMetaRefModel = (*customAsMetaRefModel)(nil)

type (
	// AsMetaRefModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMetaRefModel.
	AsMetaRefModel interface {
		asMetaRefModel
	}

	customAsMetaRefModel struct {
		*defaultAsMetaRefModel
	}
)

// NewAsMetaRefModel returns a model for the database table.
func NewAsMetaRefModel(conn sqlx.SqlConn, c cache.CacheConf) AsMetaRefModel {
	return &customAsMetaRefModel{
		defaultAsMetaRefModel: newAsMetaRefModel(conn, c),
	}
}
