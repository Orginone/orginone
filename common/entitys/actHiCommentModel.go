package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiCommentModel = (*customActHiCommentModel)(nil)

type (
	// ActHiCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiCommentModel.
	ActHiCommentModel interface {
		actHiCommentModel
	}

	customActHiCommentModel struct {
		*defaultActHiCommentModel
	}
)

// NewActHiCommentModel returns a model for the database table.
func NewActHiCommentModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiCommentModel {
	return &customActHiCommentModel{
		defaultActHiCommentModel: newActHiCommentModel(conn, c),
	}
}
