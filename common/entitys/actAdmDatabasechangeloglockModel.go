package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAdmDatabasechangeloglockModel = (*customActAdmDatabasechangeloglockModel)(nil)

type (
	// ActAdmDatabasechangeloglockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAdmDatabasechangeloglockModel.
	ActAdmDatabasechangeloglockModel interface {
		actAdmDatabasechangeloglockModel
	}

	customActAdmDatabasechangeloglockModel struct {
		*defaultActAdmDatabasechangeloglockModel
	}
)

// NewActAdmDatabasechangeloglockModel returns a model for the database table.
func NewActAdmDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) ActAdmDatabasechangeloglockModel {
	return &customActAdmDatabasechangeloglockModel{
		defaultActAdmDatabasechangeloglockModel: newActAdmDatabasechangeloglockModel(conn, c),
	}
}
