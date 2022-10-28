package entitys

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsFormValueAutofillModel = (*customAsFormValueAutofillModel)(nil)

type (
	// AsFormValueAutofillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsFormValueAutofillModel.
	AsFormValueAutofillModel interface {
		asFormValueAutofillModel
	}

	customAsFormValueAutofillModel struct {
		*defaultAsFormValueAutofillModel
	}
)

// NewAsFormValueAutofillModel returns a model for the database table.
func NewAsFormValueAutofillModel(conn sqlx.SqlConn, c cache.CacheConf) AsFormValueAutofillModel {
	return &customAsFormValueAutofillModel{
		defaultAsFormValueAutofillModel: newAsFormValueAutofillModel(conn, c),
	}
}
