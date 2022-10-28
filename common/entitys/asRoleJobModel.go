package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsRoleJobModel = (*customAsRoleJobModel)(nil)

type (
	// AsRoleJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsRoleJobModel.
	AsRoleJobModel interface {
		asRoleJobModel
	}

	customAsRoleJobModel struct {
		*defaultAsRoleJobModel
	}
)

// NewAsRoleJobModel returns a model for the database table.
func NewAsRoleJobModel(conn sqlx.SqlConn, c cache.CacheConf) AsRoleJobModel {
	return &customAsRoleJobModel{
		defaultAsRoleJobModel: newAsRoleJobModel(conn, c),
	}
}
