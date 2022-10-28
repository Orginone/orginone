package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActAdmServerConfigModel = (*customActAdmServerConfigModel)(nil)

type (
	// ActAdmServerConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActAdmServerConfigModel.
	ActAdmServerConfigModel interface {
		actAdmServerConfigModel
	}

	customActAdmServerConfigModel struct {
		*defaultActAdmServerConfigModel
	}
)

// NewActAdmServerConfigModel returns a model for the database table.
func NewActAdmServerConfigModel(conn sqlx.SqlConn, c cache.CacheConf) ActAdmServerConfigModel {
	return &customActAdmServerConfigModel{
		defaultActAdmServerConfigModel: newActAdmServerConfigModel(conn, c),
	}
}
