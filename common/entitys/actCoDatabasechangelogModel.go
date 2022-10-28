package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCoDatabasechangelogModel = (*customActCoDatabasechangelogModel)(nil)

type (
	// ActCoDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCoDatabasechangelogModel.
	ActCoDatabasechangelogModel interface {
		actCoDatabasechangelogModel
	}

	customActCoDatabasechangelogModel struct {
		*defaultActCoDatabasechangelogModel
	}
)

// NewActCoDatabasechangelogModel returns a model for the database table.
func NewActCoDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActCoDatabasechangelogModel {
	return &customActCoDatabasechangelogModel{
		defaultActCoDatabasechangelogModel: newActCoDatabasechangelogModel(conn, c),
	}
}
