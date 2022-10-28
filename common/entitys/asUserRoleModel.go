package entitys

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsUserRoleModel = (*customAsUserRoleModel)(nil)

type (
	// AsUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsUserRoleModel.
	AsUserRoleModel interface {
		asUserRoleModel
		FindByRoleIdOrUserId(ctx context.Context, roleId int64, userId int64) ([]*AsUserRole, error)
	}

	customAsUserRoleModel struct {
		*defaultAsUserRoleModel
	}
)

// NewAsUserRoleModel returns a model for the database table.
func NewAsUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) AsUserRoleModel {
	return &customAsUserRoleModel{
		defaultAsUserRoleModel: newAsUserRoleModel(conn, c),
	}
}

//根据用户Id或角色Id查找
func (m *customAsUserRoleModel) FindByRoleIdOrUserId(ctx context.Context, roleId int64, userId int64) ([]*AsUserRole, error) {
	if userId == 0 && roleId == 0 {
		return nil, errors.New("params is nil.")
	}
	var selectBuilder squirrel.SelectBuilder
	if roleId > 0 {
		selectBuilder = squirrel.Select(asUserRoleRows).From(m.table).Where("role_id = ?", roleId)
	}
	if userId > 0 {
		selectBuilder = squirrel.Select(asUserRoleRows).From(m.table).Where("user_id = ?", userId)
	}
	query, values, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*AsUserRole
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
