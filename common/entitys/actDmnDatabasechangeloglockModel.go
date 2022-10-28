package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDmnDatabasechangeloglockModel = (*customActDmnDatabasechangeloglockModel)(nil)

type (
	// ActDmnDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDmnDatabasechangeloglockModel.
	ActDmnDatabasechangeloglockModel interface {
		actDmnDatabasechangeloglockModel
	}

	customActDmnDatabasechangeloglockModel struct {
		*defaultActDmnDatabasechangeloglockModel
	}
)

// NewActDmnDatabasechangeloglockModel returns a model for the database table.
func NewActDmnDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActDmnDatabasechangeloglockModel {
	return &customActDmnDatabasechangeloglockModel{
		defaultActDmnDatabasechangeloglockModel: newActDmnDatabasechangeloglockModel(conn, c),
	}
}
