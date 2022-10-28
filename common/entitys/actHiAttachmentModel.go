package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActHiAttachmentModel = (*customActHiAttachmentModel)(nil)

type (
	// ActHiAttachmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActHiAttachmentModel.
	ActHiAttachmentModel interface {
		actHiAttachmentModel
	}

	customActHiAttachmentModel struct {
		*defaultActHiAttachmentModel
	}
)

// NewActHiAttachmentModel returns a model for the database table.
func NewActHiAttachmentModel(conn sqlx.SqlConn, c cache.CacheConf) ActHiAttachmentModel {
	return &customActHiAttachmentModel{
		defaultActHiAttachmentModel: newActHiAttachmentModel(conn, c),
	}
}
