package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwEvDatabasechangelogModel = (*customFlwEvDatabasechangelogModel)(nil)

type (
	// FlwEvDatabasechangelogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwEvDatabasechangelogModel.
	FlwEvDatabasechangelogModel interface {
		flwEvDatabasechangelogModel
	}

	customFlwEvDatabasechangelogModel struct {
		*defaultFlwEvDatabasechangelogModel
	}
)

// NewFlwEvDatabasechangelogModel returns a model for the database table.
func NewFlwEvDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) FlwEvDatabasechangelogModel {
	return &customFlwEvDatabasechangelogModel{
		defaultFlwEvDatabasechangelogModel: newFlwEvDatabasechangelogModel(conn, c),
	}
}
