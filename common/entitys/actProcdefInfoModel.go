package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActProcdefInfoModel = (*customActProcdefInfoModel)(nil)

type (
	// ActProcdefInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActProcdefInfoModel.
	ActProcdefInfoModel interface {
		actProcdefInfoModel
	}

	customActProcdefInfoModel struct {
		*defaultActProcdefInfoModel
	}
)

// NewActProcdefInfoModel returns a model for the database table.
func NewActProcdefInfoModel(conn sqlx.SqlConn, c cache.CacheConf) ActProcdefInfoModel {
	return &customActProcdefInfoModel{
		defaultActProcdefInfoModel: newActProcdefInfoModel(conn, c),
	}
}
