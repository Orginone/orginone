package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsTempletProcNodeModel = (*customAsTempletProcNodeModel)(nil)

type (
	// AsTempletProcNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsTempletProcNodeModel.
	AsTempletProcNodeModel interface {
		asTempletProcNodeModel
	}

	customAsTempletProcNodeModel struct {
		*defaultAsTempletProcNodeModel
	}
)

// NewAsTempletProcNodeModel returns a model for the database table.
func NewAsTempletProcNodeModel(conn sqlx.SqlConn, c cache.CacheConf) AsTempletProcNodeModel {
	return &customAsTempletProcNodeModel{
		defaultAsTempletProcNodeModel: newAsTempletProcNodeModel(conn, c),
	}
}
