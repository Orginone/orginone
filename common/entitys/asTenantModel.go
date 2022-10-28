package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTenantModel = (*customAsTenantModel)(nil)

type (
	// AsTenantModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTenantModel.
	AsTenantModel interface {
		asTenantModel
	}

	customAsTenantModel struct {
		*defaultAsTenantModel
	}
)

// NewAsTenantModel returns a model for the database table.
func NewAsTenantModel(conn sqlx.SqlConn, c cache.CacheConf) AsTenantModel {
	return &customAsTenantModel{
		defaultAsTenantModel: newAsTenantModel(conn, c),
	}
}
