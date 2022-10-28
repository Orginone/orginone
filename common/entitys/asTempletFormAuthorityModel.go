package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTempletFormAuthorityModel = (*customAsTempletFormAuthorityModel)(nil)

type (
	// AsTempletFormAuthorityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTempletFormAuthorityModel.
	AsTempletFormAuthorityModel interface {
		asTempletFormAuthorityModel
	}

	customAsTempletFormAuthorityModel struct {
		*defaultAsTempletFormAuthorityModel
	}
)

// NewAsTempletFormAuthorityModel returns a model for the database table.
func NewAsTempletFormAuthorityModel(conn sqlx.SqlConn, c cache.CacheConf) AsTempletFormAuthorityModel {
	return &customAsTempletFormAuthorityModel{
		defaultAsTempletFormAuthorityModel: newAsTempletFormAuthorityModel(conn, c),
	}
}
