package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwEventResourceModel = (*customFlwEventResourceModel)(nil)

type (
	// FlwEventResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwEventResourceModel.
	FlwEventResourceModel interface {
		flwEventResourceModel
	}

	customFlwEventResourceModel struct {
		*defaultFlwEventResourceModel
	}
)

// NewFlwEventResourceModel returns a model for the database table.
func NewFlwEventResourceModel(conn sqlx.SqlConn, c cache.CacheConf) FlwEventResourceModel {
	return &customFlwEventResourceModel{
		defaultFlwEventResourceModel: newFlwEventResourceModel(conn, c),
	}
}
