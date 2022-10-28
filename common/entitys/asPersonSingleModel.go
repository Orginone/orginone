package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsPersonSingleModel = (*customAsPersonSingleModel)(nil)

type (
	// AsPersonSingleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsPersonSingleModel.
	AsPersonSingleModel interface {
		asPersonSingleModel
	}

	customAsPersonSingleModel struct {
		*defaultAsPersonSingleModel
	}
)

// NewAsPersonSingleModel returns a model for the database table.
func NewAsPersonSingleModel(conn sqlx.SqlConn, c cache.CacheConf) AsPersonSingleModel {
	return &customAsPersonSingleModel{
		defaultAsPersonSingleModel: newAsPersonSingleModel(conn, c),
	}
}
