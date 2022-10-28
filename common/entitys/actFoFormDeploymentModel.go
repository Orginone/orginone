package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActFoFormDeploymentModel = (*customActFoFormDeploymentModel)(nil)

type (
	// ActFoFormDeploymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActFoFormDeploymentModel.
	ActFoFormDeploymentModel interface {
		actFoFormDeploymentModel
	}

	customActFoFormDeploymentModel struct {
		*defaultActFoFormDeploymentModel
	}
)

// NewActFoFormDeploymentModel returns a model for the database table.
func NewActFoFormDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) ActFoFormDeploymentModel {
	return &customActFoFormDeploymentModel{
		defaultActFoFormDeploymentModel: newActFoFormDeploymentModel(conn, c),
	}
}
