package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BasedataAssetclassModel = (*customBasedataAssetclassModel)(nil)

type (
	// BasedataAssetclassModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBasedataAssetclassModel.
	BasedataAssetclassModel interface {
		basedataAssetclassModel
	}

	customBasedataAssetclassModel struct {
		*defaultBasedataAssetclassModel
	}
)

// NewBasedataAssetclassModel returns a model for the database table.
func NewBasedataAssetclassModel(conn sqlx.SqlConn, c cache.CacheConf) BasedataAssetclassModel {
	return &customBasedataAssetclassModel{
		defaultBasedataAssetclassModel: newBasedataAssetclassModel(conn, c),
	}
}
