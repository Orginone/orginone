package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAppDeploymentResourceModel = (*customActAppDeploymentResourceModel)(nil)

type (
	// ActAppDeploymentResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAppDeploymentResourceModel.
	ActAppDeploymentResourceModel interface {
		actAppDeploymentResourceModel
	}

	customActAppDeploymentResourceModel struct {
		*defaultActAppDeploymentResourceModel
	}
)

// NewActAppDeploymentResourceModel returns a model for the database table.
func NewActAppDeploymentResourceModel(conn sqlx.SqlConn, c cache.CacheConf) ActAppDeploymentResourceModel {
	return &customActAppDeploymentResourceModel{
		defaultActAppDeploymentResourceModel: newActAppDeploymentResourceModel(conn, c),
	}
}
