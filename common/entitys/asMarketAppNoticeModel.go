package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppNoticeModel = (*customAsMarketAppNoticeModel)(nil)

type (
	// AsMarketAppNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppNoticeModel.
	AsMarketAppNoticeModel interface {
		asMarketAppNoticeModel
	}

	customAsMarketAppNoticeModel struct {
		*defaultAsMarketAppNoticeModel
	}
)

// NewAsMarketAppNoticeModel returns a model for the database table.
func NewAsMarketAppNoticeModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppNoticeModel {
	return &customAsMarketAppNoticeModel{
		defaultAsMarketAppNoticeModel: newAsMarketAppNoticeModel(conn, c),
	}
}
