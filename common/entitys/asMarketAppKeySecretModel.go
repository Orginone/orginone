package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsMarketAppKeySecretModel = (*customAsMarketAppKeySecretModel)(nil)

type (
	// AsMarketAppKeySecretModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsMarketAppKeySecretModel.
	AsMarketAppKeySecretModel interface {
		asMarketAppKeySecretModel
	}

	customAsMarketAppKeySecretModel struct {
		*defaultAsMarketAppKeySecretModel
	}
)

// NewAsMarketAppKeySecretModel returns a model for the database table.
func NewAsMarketAppKeySecretModel(conn sqlx.SqlConn, c cache.CacheConf) AsMarketAppKeySecretModel {
	return &customAsMarketAppKeySecretModel{
		defaultAsMarketAppKeySecretModel: newAsMarketAppKeySecretModel(conn, c),
	}
}
