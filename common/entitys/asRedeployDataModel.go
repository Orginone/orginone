package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsRedeployDataModel = (*customAsRedeployDataModel)(nil)

type (
	// AsRedeployDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsRedeployDataModel.
	AsRedeployDataModel interface {
		asRedeployDataModel
	}

	customAsRedeployDataModel struct {
		*defaultAsRedeployDataModel
	}
)

// NewAsRedeployDataModel returns a model for the database table.
func NewAsRedeployDataModel(conn sqlx.SqlConn, c cache.CacheConf) AsRedeployDataModel {
	return &customAsRedeployDataModel{
		defaultAsRedeployDataModel: newAsRedeployDataModel(conn, c),
	}
}
