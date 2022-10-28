package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BasedataAssetsortGbModel = (*customBasedataAssetsortGbModel)(nil)

type (
	// BasedataAssetsortGbModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBasedataAssetsortGbModel.
	BasedataAssetsortGbModel interface {
		basedataAssetsortGbModel
	}

	customBasedataAssetsortGbModel struct {
		*defaultBasedataAssetsortGbModel
	}
)

// NewBasedataAssetsortGbModel returns a model for the database table.
func NewBasedataAssetsortGbModel(conn sqlx.SqlConn, c cache.CacheConf) BasedataAssetsortGbModel {
	return &customBasedataAssetsortGbModel{
		defaultBasedataAssetsortGbModel: newBasedataAssetsortGbModel(conn, c),
	}
}
