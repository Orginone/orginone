package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WorkReplyModel = (*customWorkReplyModel)(nil)

type (
	// WorkReplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWorkReplyModel.
	WorkReplyModel interface {
		workReplyModel
	}

	customWorkReplyModel struct {
		*defaultWorkReplyModel
	}
)

// NewWorkReplyModel returns a model for the database table.
func NewWorkReplyModel(conn sqlx.SqlConn, c cache.CacheConf) WorkReplyModel {
	return &customWorkReplyModel{
		defaultWorkReplyModel: newWorkReplyModel(conn, c),
	}
}
