package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsJobModel = (*customAsJobModel)(nil)

type (
	// AsJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsJobModel.
	AsJobModel interface {
		asJobModel
	}

	customAsJobModel struct {
		*defaultAsJobModel
	}
)

// NewAsJobModel returns a model for the database table.
func NewAsJobModel(conn sqlx.SqlConn, c cache.CacheConf) AsJobModel {
	return &customAsJobModel{
		defaultAsJobModel: newAsJobModel(conn, c),
	}
}
