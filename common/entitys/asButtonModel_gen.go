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
	asButtonFieldNames          = builder.RawFieldNames(&AsButton{})
	asButtonRows                = strings.Join(asButtonFieldNames, ",")
	asButtonRowsExpectAutoSet   = strings.Join(stringx.Remove(asButtonFieldNames, "`create_time`", "`update_time`"), ",")
	asButtonRowsWithPlaceHolder = strings.Join(stringx.Remove(asButtonFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsButtonIdPrefix = "cache:asset:asButton:id:"
)

type (
	asButtonModel interface {
		Insert(ctx context.Context, data *AsButton) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsButton, error)
		Update(ctx context.Context, data *AsButton) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsButtonModel struct {
		sqlc.CachedConn
		table string
	}

	AsButton struct {
		Id      string         `db:"id"`       // 主键id
		Name    string         `db:"name"`     // 按钮显示的名称
		Alias   sql.NullString `db:"alias"`    // 按钮别名
		IconCls sql.NullString `db:"icon_cls"` // 按钮图标
		Action  sql.NullString `db:"action"`   // 该按钮默认绑定的访问接口
	}
)

func newAsButtonModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsButtonModel {
	return &defaultAsButtonModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_button`",
	}
}

func (m *defaultAsButtonModel) Insert(ctx context.Context, data *AsButton) (sql.Result, error) {
	assetAsButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsButtonIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, asButtonRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.Alias, data.IconCls, data.Action)
	}, assetAsButtonIdKey)
	return ret, err
}

func (m *defaultAsButtonModel) FindOne(ctx context.Context, id string) (*AsButton, error) {
	assetAsButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsButtonIdPrefix, id)
	var resp AsButton
	err := m.QueryRowCtx(ctx, &resp, assetAsButtonIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asButtonRows, m.table)
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

func (m *defaultAsButtonModel) Update(ctx context.Context, data *AsButton) error {
	assetAsButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsButtonIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asButtonRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Alias, data.IconCls, data.Action, data.Id)
	}, assetAsButtonIdKey)
	return err
}

func (m *defaultAsButtonModel) Delete(ctx context.Context, id string) error {
	assetAsButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsButtonIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsButtonIdKey)
	return err
}

func (m *defaultAsButtonModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsButtonIdPrefix, primary)
}

func (m *defaultAsButtonModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asButtonRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsButtonModel) tableName() string {
	return m.table
}
