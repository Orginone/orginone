package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeBlogModel = (*customBladeBlogModel)(nil)

type (
	// BladeBlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeBlogModel.
	BladeBlogModel interface {
		bladeBlogModel
	}

	customBladeBlogModel struct {
		*defaultBladeBlogModel
	}
)

// NewBladeBlogModel returns a model for the database table.
func NewBladeBlogModel(conn sqlx.SqlConn, c cache.CacheConf) BladeBlogModel {
	return &customBladeBlogModel{
		defaultBladeBlogModel: newBladeBlogModel(conn, c),
	}
}
