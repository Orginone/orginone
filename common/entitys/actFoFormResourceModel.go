package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActFoFormResourceModel = (*customActFoFormResourceModel)(nil)

type (
	// ActFoFormResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActFoFormResourceModel.
	ActFoFormResourceModel interface {
		actFoFormResourceModel
	}

	customActFoFormResourceModel struct {
		*defaultActFoFormResourceModel
	}
)

// NewActFoFormResourceModel returns a model for the database table.
func NewActFoFormResourceModel(conn sqlx.SqlConn, c cache.CacheConf) ActFoFormResourceModel {
	return &customActFoFormResourceModel{
		defaultActFoFormResourceModel: newActFoFormResourceModel(conn, c),
	}
}
