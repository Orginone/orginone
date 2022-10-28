package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDmnDatabasechangelogModel = (*customActDmnDatabasechangelogModel)(nil)

type (
	// ActDmnDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDmnDatabasechangelogModel.
	ActDmnDatabasechangelogModel interface {
		actDmnDatabasechangelogModel
	}

	customActDmnDatabasechangelogModel struct {
		*defaultActDmnDatabasechangelogModel
	}
)

// NewActDmnDatabasechangelogModel returns a model for the database table.
func NewActDmnDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActDmnDatabasechangelogModel {
	return &customActDmnDatabasechangelogModel{
		defaultActDmnDatabasechangelogModel: newActDmnDatabasechangelogModel(conn, c),
	}
}
