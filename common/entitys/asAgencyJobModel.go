package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsAgencyJobModel = (*customAsAgencyJobModel)(nil)

type (
	// AsAgencyJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsAgencyJobModel.
	AsAgencyJobModel interface {
		asAgencyJobModel
	}

	customAsAgencyJobModel struct {
		*defaultAsAgencyJobModel
	}
)

// NewAsAgencyJobModel returns a model for the database table.
func NewAsAgencyJobModel(conn sqlx.SqlConn, c cache.CacheConf) AsAgencyJobModel {
	return &customAsAgencyJobModel{
		defaultAsAgencyJobModel: newAsAgencyJobModel(conn, c),
	}
}
