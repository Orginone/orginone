package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsAllGroupModel = (*customAsAllGroupModel)(nil)

type (
	// AsAllGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsAllGroupModel.
	AsAllGroupModel interface {
		asAllGroupModel
	}

	customAsAllGroupModel struct {
		*defaultAsAllGroupModel
	}
)

// NewAsAllGroupModel returns a model for the database table.
func NewAsAllGroupModel(conn sqlx.SqlConn, c cache.CacheConf) AsAllGroupModel {
	return &customAsAllGroupModel{
		defaultAsAllGroupModel: newAsAllGroupModel(conn, c),
	}
}
