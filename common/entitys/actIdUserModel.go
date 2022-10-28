package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdUserModel = (*customActIdUserModel)(nil)

type (
	// ActIdUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdUserModel.
	ActIdUserModel interface {
		actIdUserModel
	}

	customActIdUserModel struct {
		*defaultActIdUserModel
	}
)

// NewActIdUserModel returns a model for the database table.
func NewActIdUserModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdUserModel {
	return &customActIdUserModel{
		defaultActIdUserModel: newActIdUserModel(conn, c),
	}
}
