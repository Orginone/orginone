package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsProcNodeModel = (*customAsProcNodeModel)(nil)

type (
	// AsProcNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsProcNodeModel.
	AsProcNodeModel interface {
		asProcNodeModel
	}

	customAsProcNodeModel struct {
		*defaultAsProcNodeModel
	}
)

// NewAsProcNodeModel returns a model for the database table.
func NewAsProcNodeModel(conn sqlx.SqlConn, c cache.CacheConf) AsProcNodeModel {
	return &customAsProcNodeModel{
		defaultAsProcNodeModel: newAsProcNodeModel(conn, c),
	}
}
