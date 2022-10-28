package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTempletApplicationModel = (*customAsTempletApplicationModel)(nil)

type (
	// AsTempletApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTempletApplicationModel.
	AsTempletApplicationModel interface {
		asTempletApplicationModel
	}

	customAsTempletApplicationModel struct {
		*defaultAsTempletApplicationModel
	}
)

// NewAsTempletApplicationModel returns a model for the database table.
func NewAsTempletApplicationModel(conn sqlx.SqlConn, c cache.CacheConf) AsTempletApplicationModel {
	return &customAsTempletApplicationModel{
		defaultAsTempletApplicationModel: newAsTempletApplicationModel(conn, c),
	}
}
