package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdInfoModel = (*customActIdInfoModel)(nil)

type (
	// ActIdInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdInfoModel.
	ActIdInfoModel interface {
		actIdInfoModel
	}

	customActIdInfoModel struct {
		*defaultActIdInfoModel
	}
)

// NewActIdInfoModel returns a model for the database table.
func NewActIdInfoModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdInfoModel {
	return &customActIdInfoModel{
		defaultActIdInfoModel: newActIdInfoModel(conn, c),
	}
}
