package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsAgencyUserModel = (*customAsAgencyUserModel)(nil)

type (
	// AsAgencyUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsAgencyUserModel.
	AsAgencyUserModel interface {
		asAgencyUserModel
	}

	customAsAgencyUserModel struct {
		*defaultAsAgencyUserModel
	}
)

// NewAsAgencyUserModel returns a model for the database table.
func NewAsAgencyUserModel(conn sqlx.SqlConn, c cache.CacheConf) AsAgencyUserModel {
	return &customAsAgencyUserModel{
		defaultAsAgencyUserModel: newAsAgencyUserModel(conn, c),
	}
}
