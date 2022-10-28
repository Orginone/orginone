// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bladeRoleMenuFieldNames          = builder.RawFieldNames(&BladeRoleMenu{})
	bladeRoleMenuRows                = strings.Join(bladeRoleMenuFieldNames, ",")
	bladeRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(bladeRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bladeRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(bladeRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBladeRoleMenuIdPrefix = "cache:asset:bladeRoleMenu:id:"
)

type (
	bladeRoleMenuModel interface {
		Insert(ctx context.Context, data *BladeRoleMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BladeRoleMenu, error)
		Update(ctx context.Context, data *BladeRoleMenu) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBladeRoleMenuModel struct {
		sqlc.CachedConn
		table string
	}

	BladeRoleMenu struct {
		Id     int64         `db:"id"`      // 主键
		MenuId sql.NullInt64 `db:"menu_id"` // 菜单id
		RoleId sql.NullInt64 `db:"role_id"` // 角色id
	}
)

func newBladeRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBladeRoleMenuModel {
	return &defaultBladeRoleMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`blade_role_menu`",
	}
}

func (m *defaultBladeRoleMenuModel) Insert(ctx context.Context, data *BladeRoleMenu) (sql.Result, error) {
	assetBladeRoleMenuIdKey := fmt.Sprintf("%s%v", cacheAssetBladeRoleMenuIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bladeRoleMenuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MenuId, data.RoleId)
	}, assetBladeRoleMenuIdKey)
	return ret, err
}

func (m *defaultBladeRoleMenuModel) FindOne(ctx context.Context, id int64) (*BladeRoleMenu, error) {
	assetBladeRoleMenuIdKey := fmt.Sprintf("%s%v", cacheAssetBladeRoleMenuIdPrefix, id)
	var resp BladeRoleMenu
	err := m.QueryRowCtx(ctx, &resp, assetBladeRoleMenuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeRoleMenuRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBladeRoleMenuModel) Update(ctx context.Context, data *BladeRoleMenu) error {
	assetBladeRoleMenuIdKey := fmt.Sprintf("%s%v", cacheAssetBladeRoleMenuIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bladeRoleMenuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MenuId, data.RoleId, data.Id)
	}, assetBladeRoleMenuIdKey)
	return err
}

func (m *defaultBladeRoleMenuModel) Delete(ctx context.Context, id int64) error {
	assetBladeRoleMenuIdKey := fmt.Sprintf("%s%v", cacheAssetBladeRoleMenuIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetBladeRoleMenuIdKey)
	return err
}

func (m *defaultBladeRoleMenuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBladeRoleMenuIdPrefix, primary)
}

func (m *defaultBladeRoleMenuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeRoleMenuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBladeRoleMenuModel) tableName() string {
	return m.table
}
