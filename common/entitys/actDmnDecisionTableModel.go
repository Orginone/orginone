package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActDmnDecisionTableModel = (*customActDmnDecisionTableModel)(nil)

type (
	// ActDmnDecisionTableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActDmnDecisionTableModel.
	ActDmnDecisionTableModel interface {
		actDmnDecisionTableModel
	}

	customActDmnDecisionTableModel struct {
		*defaultActDmnDecisionTableModel
	}
)

// NewActDmnDecisionTableModel returns a model for the database table.
func NewActDmnDecisionTableModel(conn sqlx.SqlConn, c cache.CacheConf) ActDmnDecisionTableModel {
	return &customActDmnDecisionTableModel{
		defaultActDmnDecisionTableModel: newActDmnDecisionTableModel(conn, c),
	}
}
