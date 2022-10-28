package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsCategoryModel = (*customAsCategoryModel)(nil)

type (
	// AsCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsCategoryModel.
	AsCategoryModel interface {
		asCategoryModel
	}

	customAsCategoryModel struct {
		*defaultAsCategoryModel
	}
)

// NewAsCategoryModel returns a model for the database table.
func NewAsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) AsCategoryModel {
	return &customAsCategoryModel{
		defaultAsCategoryModel: newAsCategoryModel(conn, c),
	}
}
