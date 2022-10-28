package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDeDatabasechangelogModel = (*customActDeDatabasechangelogModel)(nil)

type (
	// ActDeDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDeDatabasechangelogModel.
	ActDeDatabasechangelogModel interface {
		actDeDatabasechangelogModel
	}

	customActDeDatabasechangelogModel struct {
		*defaultActDeDatabasechangelogModel
	}
)

// NewActDeDatabasechangelogModel returns a model for the database table.
func NewActDeDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActDeDatabasechangelogModel {
	return &customActDeDatabasechangelogModel{
		defaultActDeDatabasechangelogModel: newActDeDatabasechangelogModel(conn, c),
	}
}
