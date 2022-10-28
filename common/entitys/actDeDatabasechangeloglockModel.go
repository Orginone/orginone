package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDeDatabasechangeloglockModel = (*customActDeDatabasechangeloglockModel)(nil)

type (
	// ActDeDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDeDatabasechangeloglockModel.
	ActDeDatabasechangeloglockModel interface {
		actDeDatabasechangeloglockModel
	}

	customActDeDatabasechangeloglockModel struct {
		*defaultActDeDatabasechangeloglockModel
	}
)

// NewActDeDatabasechangeloglockModel returns a model for the database table.
func NewActDeDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActDeDatabasechangeloglockModel {
	return &customActDeDatabasechangeloglockModel{
		defaultActDeDatabasechangeloglockModel: newActDeDatabasechangeloglockModel(conn, c),
	}
}
