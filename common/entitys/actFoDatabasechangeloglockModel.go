package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActFoDatabasechangeloglockModel = (*customActFoDatabasechangeloglockModel)(nil)

type (
	// ActFoDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActFoDatabasechangeloglockModel.
	ActFoDatabasechangeloglockModel interface {
		actFoDatabasechangeloglockModel
	}

	customActFoDatabasechangeloglockModel struct {
		*defaultActFoDatabasechangeloglockModel
	}
)

// NewActFoDatabasechangeloglockModel returns a model for the database table.
func NewActFoDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActFoDatabasechangeloglockModel {
	return &customActFoDatabasechangeloglockModel{
		defaultActFoDatabasechangeloglockModel: newActFoDatabasechangeloglockModel(conn, c),
	}
}
