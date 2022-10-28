package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActIdMembershipModel = (*customActIdMembershipModel)(nil)

type (
	// ActIdMembershipModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActIdMembershipModel.
	ActIdMembershipModel interface {
		actIdMembershipModel
	}

	customActIdMembershipModel struct {
		*defaultActIdMembershipModel
	}
)

// NewActIdMembershipModel returns a model for the database table.
func NewActIdMembershipModel(conn sqlx.SqlConn, c cache.CacheConf) ActIdMembershipModel {
	return &customActIdMembershipModel{
		defaultActIdMembershipModel: newActIdMembershipModel(conn, c),
	}
}
