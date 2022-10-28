package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BaseinfoAdministrativeAreaAllModel = (*customBaseinfoAdministrativeAreaAllModel)(nil)

type (
	// BaseinfoAdministrativeAreaAllModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBaseinfoAdministrativeAreaAllModel.
	BaseinfoAdministrativeAreaAllModel interface {
		baseinfoAdministrativeAreaAllModel
	}

	customBaseinfoAdministrativeAreaAllModel struct {
		*defaultBaseinfoAdministrativeAreaAllModel
	}
)

// NewBaseinfoAdministrativeAreaAllModel returns a model for the database table.
func NewBaseinfoAdministrativeAreaAllModel(conn sqlx.SqlConn, c cache.CacheConf) BaseinfoAdministrativeAreaAllModel {
	return &customBaseinfoAdministrativeAreaAllModel{
		defaultBaseinfoAdministrativeAreaAllModel: newBaseinfoAdministrativeAreaAllModel(conn, c),
	}
}
