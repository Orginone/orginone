package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTempletFormModelModel = (*customAsTempletFormModelModel)(nil)

type (
	// AsTempletFormModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTempletFormModelModel.
	AsTempletFormModelModel interface {
		asTempletFormModelModel
	}

	customAsTempletFormModelModel struct {
		*defaultAsTempletFormModelModel
	}
)

// NewAsTempletFormModelModel returns a model for the database table.
func NewAsTempletFormModelModel(conn sqlx.SqlConn, c cache.CacheConf) AsTempletFormModelModel {
	return &customAsTempletFormModelModel{
		defaultAsTempletFormModelModel: newAsTempletFormModelModel(conn, c),
	}
}
