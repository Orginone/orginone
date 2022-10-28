package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FlwEvDatabasechangeloglockModel = (*customFlwEvDatabasechangeloglockModel)(nil)

type (
	// FlwEvDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFlwEvDatabasechangeloglockModel.
	FlwEvDatabasechangeloglockModel interface {
		flwEvDatabasechangeloglockModel
	}

	customFlwEvDatabasechangeloglockModel struct {
		*defaultFlwEvDatabasechangeloglockModel
	}
)

// NewFlwEvDatabasechangeloglockModel returns a model for the database table.
func NewFlwEvDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) FlwEvDatabasechangeloglockModel {
	return &customFlwEvDatabasechangeloglockModel{
		defaultFlwEvDatabasechangeloglockModel: newFlwEvDatabasechangeloglockModel(conn, c),
	}
}
