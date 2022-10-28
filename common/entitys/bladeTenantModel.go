package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeTenantModel = (*customBladeTenantModel)(nil)

type (
	// BladeTenantModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeTenantModel.
	BladeTenantModel interface {
		bladeTenantModel
	}

	customBladeTenantModel struct {
		*defaultBladeTenantModel
	}
)

// NewBladeTenantModel returns a model for the database table.
func NewBladeTenantModel(conn sqlx.SqlConn, c cache.CacheConf) BladeTenantModel {
	return &customBladeTenantModel{
		defaultBladeTenantModel: newBladeTenantModel(conn, c),
	}
}
