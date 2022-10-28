package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserJobModel = (*customAsUserJobModel)(nil)

type (
	// AsUserJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserJobModel.
	AsUserJobModel interface {
		asUserJobModel
	}

	customAsUserJobModel struct {
		*defaultAsUserJobModel
	}
)

// NewAsUserJobModel returns a model for the database table.
func NewAsUserJobModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserJobModel {
	return &customAsUserJobModel{
		defaultAsUserJobModel: newAsUserJobModel(conn, c),
	}
}
