package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnDeploymentModel = (*customActCmmnDeploymentModel)(nil)

type (
	// ActCmmnDeploymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnDeploymentModel.
	ActCmmnDeploymentModel interface {
		actCmmnDeploymentModel
	}

	customActCmmnDeploymentModel struct {
		*defaultActCmmnDeploymentModel
	}
)

// NewActCmmnDeploymentModel returns a model for the database table.
func NewActCmmnDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnDeploymentModel {
	return &customActCmmnDeploymentModel{
		defaultActCmmnDeploymentModel: newActCmmnDeploymentModel(conn, c),
	}
}
