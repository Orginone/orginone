package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAdmDatabasechangelogModel = (*customActAdmDatabasechangelogModel)(nil)

type (
	// ActAdmDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAdmDatabasechangelogModel.
	ActAdmDatabasechangelogModel interface {
		actAdmDatabasechangelogModel
	}

	customActAdmDatabasechangelogModel struct {
		*defaultActAdmDatabasechangelogModel
	}
)

// NewActAdmDatabasechangelogModel returns a model for the database table.
func NewActAdmDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActAdmDatabasechangelogModel {
	return &customActAdmDatabasechangelogModel{
		defaultActAdmDatabasechangelogModel: newActAdmDatabasechangelogModel(conn, c),
	}
}
