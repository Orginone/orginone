package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCmmnDatabasechangeloglockModel = (*customActCmmnDatabasechangeloglockModel)(nil)

type (
	// ActCmmnDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCmmnDatabasechangeloglockModel.
	ActCmmnDatabasechangeloglockModel interface {
		actCmmnDatabasechangeloglockModel
	}

	customActCmmnDatabasechangeloglockModel struct {
		*defaultActCmmnDatabasechangeloglockModel
	}
)

// NewActCmmnDatabasechangeloglockModel returns a model for the database table.
func NewActCmmnDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActCmmnDatabasechangeloglockModel {
	return &customActCmmnDatabasechangeloglockModel{
		defaultActCmmnDatabasechangeloglockModel: newActCmmnDatabasechangeloglockModel(conn, c),
	}
}
