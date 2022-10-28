package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsProcInstModel = (*customAsProcInstModel)(nil)

type (
	// AsProcInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsProcInstModel.
	AsProcInstModel interface {
		asProcInstModel
	}

	customAsProcInstModel struct {
		*defaultAsProcInstModel
	}
)

// NewAsProcInstModel returns a model for the database table.
func NewAsProcInstModel(conn sqlx.SqlConn, c cache.CacheConf) AsProcInstModel {
	return &customAsProcInstModel{
		defaultAsProcInstModel: newAsProcInstModel(conn, c),
	}
}
