package entitys

import (
	"context"
	"fmt"
	"orginone/common/tools"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsRoleModel = (*customAsRoleModel)(nil)

type (
	// AsRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsRoleModel.
	AsRoleModel interface {
		asRoleModel
		FindByRoleIds(ctx context.Context, roleIds []int64) ([]*AsRole, error)
	}

	customAsRoleModel struct {
		*defaultAsRoleModel
	}
)

// NewAsRoleModel returns a model for the database table.
func NewAsRoleModel(conn sqlx.SqlConn, c cache.CacheConf) AsRoleModel {
	return &customAsRoleModel{
		defaultAsRoleModel: newAsRoleModel(conn, c),
	}
}

//根据角色Ids查找
func (m *customAsRoleModel) FindByRoleIds(ctx context.Context, roleIds []int64) ([]*AsRole, error) {
	query, values, err := squirrel.Select(asRoleRows).From(m.table).Where(fmt.Sprintf("id in %s", tools.SqlInArgsMap(roleIds))).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*AsRole
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
