package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTempletDeModelModel = (*customAsTempletDeModelModel)(nil)

type (
	// AsTempletDeModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTempletDeModelModel.
	AsTempletDeModelModel interface {
		asTempletDeModelModel
	}

	customAsTempletDeModelModel struct {
		*defaultAsTempletDeModelModel
	}
)

// NewAsTempletDeModelModel returns a model for the database table.
func NewAsTempletDeModelModel(conn sqlx.SqlConn, c cache.CacheConf) AsTempletDeModelModel {
	return &customAsTempletDeModelModel{
		defaultAsTempletDeModelModel: newAsTempletDeModelModel(conn, c),
	}
}
