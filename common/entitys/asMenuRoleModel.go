package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMenuRoleModel = (*customAsMenuRoleModel)(nil)

type (
	// AsMenuRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMenuRoleModel.
	AsMenuRoleModel interface {
		asMenuRoleModel
	}

	customAsMenuRoleModel struct {
		*defaultAsMenuRoleModel
	}
)

// NewAsMenuRoleModel returns a model for the database table.
func NewAsMenuRoleModel(conn sqlx.SqlConn, c cache.CacheConf) AsMenuRoleModel {
	return &customAsMenuRoleModel{
		defaultAsMenuRoleModel: newAsMenuRoleModel(conn, c),
	}
}
