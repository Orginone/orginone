package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDmnDeploymentResourceModel = (*customActDmnDeploymentResourceModel)(nil)

type (
	// ActDmnDeploymentResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDmnDeploymentResourceModel.
	ActDmnDeploymentResourceModel interface {
		actDmnDeploymentResourceModel
	}

	customActDmnDeploymentResourceModel struct {
		*defaultActDmnDeploymentResourceModel
	}
)

// NewActDmnDeploymentResourceModel returns a model for the database table.
func NewActDmnDeploymentResourceModel(conn sqlx.SqlConn, c cache.CacheConf) ActDmnDeploymentResourceModel {
	return &customActDmnDeploymentResourceModel{
		defaultActDmnDeploymentResourceModel: newActDmnDeploymentResourceModel(conn, c),
	}
}
