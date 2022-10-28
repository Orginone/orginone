package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActFoFormInstanceModel = (*customActFoFormInstanceModel)(nil)

type (
	// ActFoFormInstanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActFoFormInstanceModel.
	ActFoFormInstanceModel interface {
		actFoFormInstanceModel
	}

	customActFoFormInstanceModel struct {
		*defaultActFoFormInstanceModel
	}
)

// NewActFoFormInstanceModel returns a model for the database table.
func NewActFoFormInstanceModel(conn sqlx.SqlConn, c cache.CacheConf) ActFoFormInstanceModel {
	return &customActFoFormInstanceModel{
		defaultActFoFormInstanceModel: newActFoFormInstanceModel(conn, c),
	}
}
