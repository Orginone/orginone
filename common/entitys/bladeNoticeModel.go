package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BladeNoticeModel = (*customBladeNoticeModel)(nil)

type (
	// BladeNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBladeNoticeModel.
	BladeNoticeModel interface {
		bladeNoticeModel
	}

	customBladeNoticeModel struct {
		*defaultBladeNoticeModel
	}
)

// NewBladeNoticeModel returns a model for the database table.
func NewBladeNoticeModel(conn sqlx.SqlConn, c cache.CacheConf) BladeNoticeModel {
	return &customBladeNoticeModel{
		defaultBladeNoticeModel: newBladeNoticeModel(conn, c),
	}
}
