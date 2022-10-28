package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormModelHistoryModel = (*customAsFormModelHistoryModel)(nil)

type (
	// AsFormModelHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormModelHistoryModel.
	AsFormModelHistoryModel interface {
		asFormModelHistoryModel
	}

	customAsFormModelHistoryModel struct {
		*defaultAsFormModelHistoryModel
	}
)

// NewAsFormModelHistoryModel returns a model for the database table.
func NewAsFormModelHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormModelHistoryModel {
	return &customAsFormModelHistoryModel{
		defaultAsFormModelHistoryModel: newAsFormModelHistoryModel(conn, c),
	}
}
