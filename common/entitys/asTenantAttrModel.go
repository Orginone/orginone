package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTenantAttrModel = (*customAsTenantAttrModel)(nil)

type (
	// AsTenantAttrModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTenantAttrModel.
	AsTenantAttrModel interface {
		asTenantAttrModel
	}

	customAsTenantAttrModel struct {
		*defaultAsTenantAttrModel
	}
)

// NewAsTenantAttrModel returns a model for the database table.
func NewAsTenantAttrModel(conn sqlx.SqlConn, c cache.CacheConf) AsTenantAttrModel {
	return &customAsTenantAttrModel{
		defaultAsTenantAttrModel: newAsTenantAttrModel(conn, c),
	}
}
