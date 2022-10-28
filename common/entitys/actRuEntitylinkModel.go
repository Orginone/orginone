package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuEntitylinkModel = (*customActRuEntitylinkModel)(nil)

type (
	// ActRuEntitylinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuEntitylinkModel.
	ActRuEntitylinkModel interface {
		actRuEntitylinkModel
	}

	customActRuEntitylinkModel struct {
		*defaultActRuEntitylinkModel
	}
)

// NewActRuEntitylinkModel returns a model for the database table.
func NewActRuEntitylinkModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuEntitylinkModel {
	return &customActRuEntitylinkModel{
		defaultActRuEntitylinkModel: newActRuEntitylinkModel(conn, c),
	}
}
