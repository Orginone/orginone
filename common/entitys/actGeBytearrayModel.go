package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActGeBytearrayModel = (*customActGeBytearrayModel)(nil)

type (
	// ActGeBytearrayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActGeBytearrayModel.
	ActGeBytearrayModel interface {
		actGeBytearrayModel
	}

	customActGeBytearrayModel struct {
		*defaultActGeBytearrayModel
	}
)

// NewActGeBytearrayModel returns a model for the database table.
func NewActGeBytearrayModel(conn sqlx.SqlConn, c cache.CacheConf) ActGeBytearrayModel {
	return &customActGeBytearrayModel{
		defaultActGeBytearrayModel: newActGeBytearrayModel(conn, c),
	}
}
