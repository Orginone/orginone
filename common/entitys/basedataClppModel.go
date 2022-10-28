package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BasedataClppModel = (*customBasedataClppModel)(nil)

type (
	// BasedataClppModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBasedataClppModel.
	BasedataClppModel interface {
		basedataClppModel
	}

	customBasedataClppModel struct {
		*defaultBasedataClppModel
	}
)

// NewBasedataClppModel returns a model for the database table.
func NewBasedataClppModel(conn sqlx.SqlConn, c cache.CacheConf) BasedataClppModel {
	return &customBasedataClppModel{
		defaultBasedataClppModel: newBasedataClppModel(conn, c),
	}
}
