package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActReDeploymentModel = (*customActReDeploymentModel)(nil)

type (
	// ActReDeploymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActReDeploymentModel.
	ActReDeploymentModel interface {
		actReDeploymentModel
	}

	customActReDeploymentModel struct {
		*defaultActReDeploymentModel
	}
)

// NewActReDeploymentModel returns a model for the database table.
func NewActReDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) ActReDeploymentModel {
	return &customActReDeploymentModel{
		defaultActReDeploymentModel: newActReDeploymentModel(conn, c),
	}
}
