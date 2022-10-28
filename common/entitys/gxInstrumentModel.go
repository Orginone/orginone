package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GxInstrumentModel = (*customGxInstrumentModel)(nil)

type (
	// GxInstrumentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGxInstrumentModel.
	GxInstrumentModel interface {
		gxInstrumentModel
	}

	customGxInstrumentModel struct {
		*defaultGxInstrumentModel
	}
)

// NewGxInstrumentModel returns a model for the database table.
func NewGxInstrumentModel(conn sqlx.SqlConn, c cache.CacheConf) GxInstrumentModel {
	return &customGxInstrumentModel{
		defaultGxInstrumentModel: newGxInstrumentModel(conn, c),
	}
}
