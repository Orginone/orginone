package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsRoleGroupModel = (*customAsRoleGroupModel)(nil)

type (
	// AsRoleGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsRoleGroupModel.
	AsRoleGroupModel interface {
		asRoleGroupModel
	}

	customAsRoleGroupModel struct {
		*defaultAsRoleGroupModel
	}
)

// NewAsRoleGroupModel returns a model for the database table.
func NewAsRoleGroupModel(conn sqlx.SqlConn, c cache.CacheConf) AsRoleGroupModel {
	return &customAsRoleGroupModel{
		defaultAsRoleGroupModel: newAsRoleGroupModel(conn, c),
	}
}
