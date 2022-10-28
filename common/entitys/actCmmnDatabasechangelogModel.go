package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnDatabasechangelogModel = (*customActCmmnDatabasechangelogModel)(nil)

type (
	// ActCmmnDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnDatabasechangelogModel.
	ActCmmnDatabasechangelogModel interface {
		actCmmnDatabasechangelogModel
	}

	customActCmmnDatabasechangelogModel struct {
		*defaultActCmmnDatabasechangelogModel
	}
)

// NewActCmmnDatabasechangelogModel returns a model for the database table.
func NewActCmmnDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnDatabasechangelogModel {
	return &customActCmmnDatabasechangelogModel{
		defaultActCmmnDatabasechangelogModel: newActCmmnDatabasechangelogModel(conn, c),
	}
}
