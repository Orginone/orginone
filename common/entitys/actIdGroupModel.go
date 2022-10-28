package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdGroupModel = (*customActIdGroupModel)(nil)

type (
	// ActIdGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdGroupModel.
	ActIdGroupModel interface {
		actIdGroupModel
	}

	customActIdGroupModel struct {
		*defaultActIdGroupModel
	}
)

// NewActIdGroupModel returns a model for the database table.
func NewActIdGroupModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdGroupModel {
	return &customActIdGroupModel{
		defaultActIdGroupModel: newActIdGroupModel(conn, c),
	}
}
