package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTempletModel = (*customAsTempletModel)(nil)

type (
	// AsTempletModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTempletModel.
	AsTempletModel interface {
		asTempletModel
	}

	customAsTempletModel struct {
		*defaultAsTempletModel
	}
)

// NewAsTempletModel returns a model for the database table.
func NewAsTempletModel(conn sqlx.SqlConn, c cache.CacheConf) AsTempletModel {
	return &customAsTempletModel{
		defaultAsTempletModel: newAsTempletModel(conn, c),
	}
}
