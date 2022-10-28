package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiEntitylinkModel = (*customActHiEntitylinkModel)(nil)

type (
	// ActHiEntitylinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiEntitylinkModel.
	ActHiEntitylinkModel interface {
		actHiEntitylinkModel
	}

	customActHiEntitylinkModel struct {
		*defaultActHiEntitylinkModel
	}
)

// NewActHiEntitylinkModel returns a model for the database table.
func NewActHiEntitylinkModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiEntitylinkModel {
	return &customActHiEntitylinkModel{
		defaultActHiEntitylinkModel: newActHiEntitylinkModel(conn, c),
	}
}
