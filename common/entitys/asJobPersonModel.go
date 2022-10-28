package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsJobPersonModel = (*customAsJobPersonModel)(nil)

type (
	// AsJobPersonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsJobPersonModel.
	AsJobPersonModel interface {
		asJobPersonModel
	}

	customAsJobPersonModel struct {
		*defaultAsJobPersonModel
	}
)

// NewAsJobPersonModel returns a model for the database table.
func NewAsJobPersonModel(conn sqlx.SqlConn, c cache.CacheConf) AsJobPersonModel {
	return &customAsJobPersonModel{
		defaultAsJobPersonModel: newAsJobPersonModel(conn, c),
	}
}
