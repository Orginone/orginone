package entitys

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsAgencyPersonModel = (*customAsAgencyPersonModel)(nil)

type (
	// AsAgencyPersonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsAgencyPersonModel.
	AsAgencyPersonModel interface {
		asAgencyPersonModel
		FindByAgencyIdOrPersonId(ctx context.Context, personId int64, agencyId int64) ([]*AsAgencyPerson, error)
	}

	customAsAgencyPersonModel struct {
		*defaultAsAgencyPersonModel
	}
)

// NewAsAgencyPersonModel returns a model for the database table.
func NewAsAgencyPersonModel(conn sqlx.SqlConn, c cache.CacheConf) AsAgencyPersonModel {
	return &customAsAgencyPersonModel{
		defaultAsAgencyPersonModel: newAsAgencyPersonModel(conn, c),
	}
}

//根据用户Id或角色Id查找
func (m *customAsAgencyPersonModel) FindByAgencyIdOrPersonId(ctx context.Context, personId int64, agencyId int64) ([]*AsAgencyPerson, error) {
	if personId == 0 && agencyId == 0 {
		return nil, errors.New("params is nil.")
	}
	var selectBuilder squirrel.SelectBuilder
	if personId > 0 {
		selectBuilder = squirrel.Select(asAgencyPersonRows).From(m.table).Where("person_id = ?", personId)
	}
	if agencyId > 0 {
		selectBuilder = squirrel.Select(asAgencyPersonRows).From(m.table).Where("agency_id = ?", agencyId)
	}
	query, values, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*AsAgencyPerson
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
