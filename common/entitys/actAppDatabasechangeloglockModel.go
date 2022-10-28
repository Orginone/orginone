package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAppDatabasechangeloglockModel = (*customActAppDatabasechangeloglockModel)(nil)

type (
	// ActAppDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAppDatabasechangeloglockModel.
	ActAppDatabasechangeloglockModel interface {
		actAppDatabasechangeloglockModel
	}

	customActAppDatabasechangeloglockModel struct {
		*defaultActAppDatabasechangeloglockModel
	}
)

// NewActAppDatabasechangeloglockModel returns a model for the database table.
func NewActAppDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActAppDatabasechangeloglockModel {
	return &customActAppDatabasechangeloglockModel{
		defaultActAppDatabasechangeloglockModel: newActAppDatabasechangeloglockModel(conn, c),
	}
}
