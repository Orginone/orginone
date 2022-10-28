package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdTokenModel = (*customActIdTokenModel)(nil)

type (
	// ActIdTokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdTokenModel.
	ActIdTokenModel interface {
		actIdTokenModel
	}

	customActIdTokenModel struct {
		*defaultActIdTokenModel
	}
)

// NewActIdTokenModel returns a model for the database table.
func NewActIdTokenModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdTokenModel {
	return &customActIdTokenModel{
		defaultActIdTokenModel: newActIdTokenModel(conn, c),
	}
}
