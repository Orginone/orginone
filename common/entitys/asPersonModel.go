package entitys

import (
	"context"
	"fmt"
	"orginone/common/tools"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsPersonModel = (*customAsPersonModel)(nil)

type (
	// AsPersonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsPersonModel.
	AsPersonModel interface {
		asPersonModel
		FindByPersonIds(ctx context.Context, personIds []int64) ([]*AsPerson, error)
		FindByUserInfo(ctx context.Context, userId int64, tenantCode string) ([]*AsPerson, error)
	}

	customAsPersonModel struct {
		*defaultAsPersonModel
	}
)

// NewAsPersonModel returns a model for the database table.
func NewAsPersonModel(conn sqlx.SqlConn, c cache.CacheConf) AsPersonModel {
	return &customAsPersonModel{
		defaultAsPersonModel: newAsPersonModel(conn, c),
	}
}

//根据用户信息查询人员信息
func (m *customAsPersonModel) FindByUserInfo(ctx context.Context, userId int64, tenantCode string) ([]*AsPerson, error) {
	query, values, err := squirrel.Select(asPersonRows).From(m.table).Where("user_id = ? and tenant_code = ?", userId, tenantCode).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*AsPerson
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//根据人员Ids查找
func (m *customAsPersonModel) FindByPersonIds(ctx context.Context, personIds []int64) ([]*AsPerson, error) {
	query, values, err := squirrel.Select(asInnerAgencyRows).From(m.table).Where(fmt.Sprintf("id in %s", tools.SqlInArgsMap(personIds))).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*AsPerson
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
