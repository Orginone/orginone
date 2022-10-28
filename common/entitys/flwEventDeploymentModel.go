package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwEventDeploymentModel = (*customFlwEventDeploymentModel)(nil)

type (
	// FlwEventDeploymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwEventDeploymentModel.
	FlwEventDeploymentModel interface {
		flwEventDeploymentModel
	}

	customFlwEventDeploymentModel struct {
		*defaultFlwEventDeploymentModel
	}
)

// NewFlwEventDeploymentModel returns a model for the database table.
func NewFlwEventDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) FlwEventDeploymentModel {
	return &customFlwEventDeploymentModel{
		defaultFlwEventDeploymentModel: newFlwEventDeploymentModel(conn, c),
	}
}
