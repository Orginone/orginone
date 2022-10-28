package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdBytearrayModel = (*customActIdBytearrayModel)(nil)

type (
	// ActIdBytearrayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdBytearrayModel.
	ActIdBytearrayModel interface {
		actIdBytearrayModel
	}

	customActIdBytearrayModel struct {
		*defaultActIdBytearrayModel
	}
)

// NewActIdBytearrayModel returns a model for the database table.
func NewActIdBytearrayModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdBytearrayModel {
	return &customActIdBytearrayModel{
		defaultActIdBytearrayModel: newActIdBytearrayModel(conn, c),
	}
}
