package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActFoDatabasechangelogModel = (*customActFoDatabasechangelogModel)(nil)

type (
	// ActFoDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActFoDatabasechangelogModel.
	ActFoDatabasechangelogModel interface {
		actFoDatabasechangelogModel
	}

	customActFoDatabasechangelogModel struct {
		*defaultActFoDatabasechangelogModel
	}
)

// NewActFoDatabasechangelogModel returns a model for the database table.
func NewActFoDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) ActFoDatabasechangelogModel {
	return &customActFoDatabasechangelogModel{
		defaultActFoDatabasechangelogModel: newActFoDatabasechangelogModel(conn, c),
	}
}
