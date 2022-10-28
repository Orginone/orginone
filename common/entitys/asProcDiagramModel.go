package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsProcDiagramModel = (*customAsProcDiagramModel)(nil)

type (
	// AsProcDiagramModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsProcDiagramModel.
	AsProcDiagramModel interface {
		asProcDiagramModel
	}

	customAsProcDiagramModel struct {
		*defaultAsProcDiagramModel
	}
)

// NewAsProcDiagramModel returns a model for the database table.
func NewAsProcDiagramModel(conn sqlx.SqlConn, c cache.CacheConf) AsProcDiagramModel {
	return &customAsProcDiagramModel{
		defaultAsProcDiagramModel: newAsProcDiagramModel(conn, c),
	}
}
