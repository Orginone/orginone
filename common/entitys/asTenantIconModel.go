package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTenantIconModel = (*customAsTenantIconModel)(nil)

type (
	// AsTenantIconModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTenantIconModel.
	AsTenantIconModel interface {
		asTenantIconModel
	}

	customAsTenantIconModel struct {
		*defaultAsTenantIconModel
	}
)

// NewAsTenantIconModel returns a model for the database table.
func NewAsTenantIconModel(conn sqlx.SqlConn, c cache.CacheConf) AsTenantIconModel {
	return &customAsTenantIconModel{
		defaultAsTenantIconModel: newAsTenantIconModel(conn, c),
	}
}
