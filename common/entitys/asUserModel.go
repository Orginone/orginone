package entitys

import (
	"context"
	"fmt"
	"orginone/common/tools"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserModel = (*customAsUserModel)(nil)

type (
	// AsUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserModel.
	AsUserModel interface {
		asUserModel
		FindByAccount(ctx context.Context, account string) ([]*AsUser, error)
		FindByUserIds(ctx context.Context, userIds []int64) ([]*AsUser, error)
	}

	customAsUserModel struct {
		*defaultAsUserModel
	}
)

// NewAsUserModel returns a model for the database table.
func NewAsUserModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserModel {
	return &customAsUserModel{
		defaultAsUserModel: newAsUserModel(conn, c),
	}
}

//根据账户查询用户
func (m *customAsUserModel) FindByAccount(ctx context.Context, account string) ([]*AsUser, error) {
	query, values, err := squirrel.Select(asUserRows).From(m.table).Where("phone_number = ? and is_deleted = 0", account).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*AsUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//根据用户Ids查找
func (m *customAsUserModel) FindByUserIds(ctx context.Context, userIds []int64) ([]*AsUser, error) {
	query, values, err := squirrel.Select(asUserRows).From(m.table).Where(fmt.Sprintf("id in %s", tools.SqlInArgsMap(userIds))).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*AsUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
