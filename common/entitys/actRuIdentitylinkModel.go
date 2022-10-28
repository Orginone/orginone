package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActRuIdentitylinkModel = (*customActRuIdentitylinkModel)(nil)

type (
	// ActRuIdentitylinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActRuIdentitylinkModel.
	ActRuIdentitylinkModel interface {
		actRuIdentitylinkModel
	}

	customActRuIdentitylinkModel struct {
		*defaultActRuIdentitylinkModel
	}
)

// NewActRuIdentitylinkModel returns a model for the database table.
func NewActRuIdentitylinkModel(conn sqlx.SqlConn, c cache.CacheConf) ActRuIdentitylinkModel {
	return &customActRuIdentitylinkModel{
		defaultActRuIdentitylinkModel: newActRuIdentitylinkModel(conn, c),
	}
}
