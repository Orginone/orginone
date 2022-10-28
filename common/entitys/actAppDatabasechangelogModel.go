package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAppDatabasechangelogModel = (*customActAppDatabasechangelogModel)(nil)

type (
	// ActAppDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAppDatabasechangelogModel.
	ActAppDatabasechangelogModel interface {
		actAppDatabasechangelogModel
	}

	customActAppDatabasechangelogModel struct {
		*defaultActAppDatabasechangelogModel
	}
)

// NewActAppDatabasechangelogModel returns a model for the database table.
func NewActAppDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActAppDatabasechangelogModel {
	return &customActAppDatabasechangelogModel{
		defaultActAppDatabasechangelogModel: newActAppDatabasechangelogModel(conn, c),
	}
}
