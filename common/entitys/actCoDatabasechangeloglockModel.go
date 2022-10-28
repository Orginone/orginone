package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActCoDatabasechangeloglockModel = (*customActCoDatabasechangeloglockModel)(nil)

type (
	// ActCoDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActCoDatabasechangeloglockModel.
	ActCoDatabasechangeloglockModel interface {
		actCoDatabasechangeloglockModel
	}

	customActCoDatabasechangeloglockModel struct {
		*defaultActCoDatabasechangeloglockModel
	}
)

// NewActCoDatabasechangeloglockModel returns a model for the database table.
func NewActCoDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActCoDatabasechangeloglockModel {
	return &customActCoDatabasechangeloglockModel{
		defaultActCoDatabasechangeloglockModel: newActCoDatabasechangeloglockModel(conn, c),
	}
}
