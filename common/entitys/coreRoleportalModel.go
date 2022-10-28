package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CoreRoleportalModel = (*customCoreRoleportalModel)(nil)

type (
	// CoreRoleportalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoreRoleportalModel.
	CoreRoleportalModel interface {
		coreRoleportalModel
	}

	customCoreRoleportalModel struct {
		*defaultCoreRoleportalModel
	}
)

// NewCoreRoleportalModel returns a model for the database table.
func NewCoreRoleportalModel(conn sqlx.SqlConn, c cache.CacheConf) CoreRoleportalModel {
	return &customCoreRoleportalModel{
		defaultCoreRoleportalModel: newCoreRoleportalModel(conn, c),
	}
}
