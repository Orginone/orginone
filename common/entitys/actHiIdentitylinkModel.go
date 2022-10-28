package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiIdentitylinkModel = (*customActHiIdentitylinkModel)(nil)

type (
	// ActHiIdentitylinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiIdentitylinkModel.
	ActHiIdentitylinkModel interface {
		actHiIdentitylinkModel
	}

	customActHiIdentitylinkModel struct {
		*defaultActHiIdentitylinkModel
	}
)

// NewActHiIdentitylinkModel returns a model for the database table.
func NewActHiIdentitylinkModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiIdentitylinkModel {
	return &customActHiIdentitylinkModel{
		defaultActHiIdentitylinkModel: newActHiIdentitylinkModel(conn, c),
	}
}
