package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsItopsDetailModel = (*customAsItopsDetailModel)(nil)

type (
	// AsItopsDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsItopsDetailModel.
	AsItopsDetailModel interface {
		asItopsDetailModel
	}

	customAsItopsDetailModel struct {
		*defaultAsItopsDetailModel
	}
)

// NewAsItopsDetailModel returns a model for the database table.
func NewAsItopsDetailModel(conn sqlx.SqlConn, c cache.CacheConf) AsItopsDetailModel {
	return &customAsItopsDetailModel{
		defaultAsItopsDetailModel: newAsItopsDetailModel(conn, c),
	}
}
