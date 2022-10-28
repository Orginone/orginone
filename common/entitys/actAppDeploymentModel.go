package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAppDeploymentModel = (*customActAppDeploymentModel)(nil)

type (
	// ActAppDeploymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAppDeploymentModel.
	ActAppDeploymentModel interface {
		actAppDeploymentModel
	}

	customActAppDeploymentModel struct {
		*defaultActAppDeploymentModel
	}
)

// NewActAppDeploymentModel returns a model for the database table.
func NewActAppDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) ActAppDeploymentModel {
	return &customActAppDeploymentModel{
		defaultActAppDeploymentModel: newActAppDeploymentModel(conn, c),
	}
}
