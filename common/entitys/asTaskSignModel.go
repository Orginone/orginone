package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTaskSignModel = (*customAsTaskSignModel)(nil)

type (
	// AsTaskSignModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTaskSignModel.
	AsTaskSignModel interface {
		asTaskSignModel
	}

	customAsTaskSignModel struct {
		*defaultAsTaskSignModel
	}
)

// NewAsTaskSignModel returns a model for the database table.
func NewAsTaskSignModel(conn sqlx.SqlConn, c cache.CacheConf) AsTaskSignModel {
	return &customAsTaskSignModel{
		defaultAsTaskSignModel: newAsTaskSignModel(conn, c),
	}
}
