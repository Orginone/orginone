package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormAuthorityModel = (*customAsFormAuthorityModel)(nil)

type (
	// AsFormAuthorityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormAuthorityModel.
	AsFormAuthorityModel interface {
		asFormAuthorityModel
	}

	customAsFormAuthorityModel struct {
		*defaultAsFormAuthorityModel
	}
)

// NewAsFormAuthorityModel returns a model for the database table.
func NewAsFormAuthorityModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormAuthorityModel {
	return &customAsFormAuthorityModel{
		defaultAsFormAuthorityModel: newAsFormAuthorityModel(conn, c),
	}
}
