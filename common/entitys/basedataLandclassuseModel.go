package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BasedataLandclassuseModel = (*customBasedataLandclassuseModel)(nil)

type (
	// BasedataLandclassuseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBasedataLandclassuseModel.
	BasedataLandclassuseModel interface {
		basedataLandclassuseModel
	}

	customBasedataLandclassuseModel struct {
		*defaultBasedataLandclassuseModel
	}
)

// NewBasedataLandclassuseModel returns a model for the database table.
func NewBasedataLandclassuseModel(conn sqlx.SqlConn, c cache.CacheConf) BasedataLandclassuseModel {
	return &customBasedataLandclassuseModel{
		defaultBasedataLandclassuseModel: newBasedataLandclassuseModel(conn, c),
	}
}
