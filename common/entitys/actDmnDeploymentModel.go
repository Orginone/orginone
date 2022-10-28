package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDmnDeploymentModel = (*customActDmnDeploymentModel)(nil)

type (
	// ActDmnDeploymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDmnDeploymentModel.
	ActDmnDeploymentModel interface {
		actDmnDeploymentModel
	}

	customActDmnDeploymentModel struct {
		*defaultActDmnDeploymentModel
	}
)

// NewActDmnDeploymentModel returns a model for the database table.
func NewActDmnDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) ActDmnDeploymentModel {
	return &customActDmnDeploymentModel{
		defaultActDmnDeploymentModel: newActDmnDeploymentModel(conn, c),
	}
}
