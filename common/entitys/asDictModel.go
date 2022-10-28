package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsDictModel = (*customAsDictModel)(nil)

type (
	// AsDictModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsDictModel.
	AsDictModel interface {
		asDictModel
	}

	customAsDictModel struct {
		*defaultAsDictModel
	}
)

// NewAsDictModel returns a model for the database table.
func NewAsDictModel(conn sqlx.SqlConn, c cache.CacheConf) AsDictModel {
	return &customAsDictModel{
		defaultAsDictModel: newAsDictModel(conn, c),
	}
}
