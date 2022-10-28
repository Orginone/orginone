package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnDeploymentResourceModel = (*customActCmmnDeploymentResourceModel)(nil)

type (
	// ActCmmnDeploymentResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnDeploymentResourceModel.
	ActCmmnDeploymentResourceModel interface {
		actCmmnDeploymentResourceModel
	}

	customActCmmnDeploymentResourceModel struct {
		*defaultActCmmnDeploymentResourceModel
	}
)

// NewActCmmnDeploymentResourceModel returns a model for the database table.
func NewActCmmnDeploymentResourceModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnDeploymentResourceModel {
	return &customActCmmnDeploymentResourceModel{
		defaultActCmmnDeploymentResourceModel: newActCmmnDeploymentResourceModel(conn, c),
	}
}
