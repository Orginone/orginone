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
	asTempletFormModelFieldNames          = builder.RawFieldNames(&AsTempletFormModel{})
	asTempletFormModelRows                = strings.Join(asTempletFormModelFieldNames, ",")
	asTempletFormModelRowsExpectAutoSet   = strings.Join(stringx.Remove(asTempletFormModelFieldNames, "`create_time`", "`update_time`"), ",")
	asTempletFormModelRowsWithPlaceHolder = strings.Join(stringx.Remove(asTempletFormModelFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsTempletFormModelIdPrefix = "cache:asset:asTempletFormModel:id:"
)

type (
	asTempletFormModelModel interface {
		Insert(ctx context.Context, data *AsTempletFormModel) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsTempletFormModel, error)
		Update(ctx context.Context, data *AsTempletFormModel) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsTempletFormModelModel struct {
		sqlc.CachedConn
		table string
	}

	AsTempletFormModel struct {
		Id         string         `db:"id"` // 表单模型模板Id
		FormName   string         `db:"form_name"`
		ModelSheet string         `db:"model_sheet"`
		IconCls    sql.NullString `db:"icon_cls"` // 表单显示的vue图标
	}
)

func newAsTempletFormModelModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsTempletFormModelModel {
	return &defaultAsTempletFormModelModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_templet_form_model`",
	}
}

func (m *defaultAsTempletFormModelModel) Insert(ctx context.Context, data *AsTempletFormModel) (sql.Result, error) {
	assetAsTempletFormModelIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormModelIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, asTempletFormModelRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.FormName, data.ModelSheet, data.IconCls)
	}, assetAsTempletFormModelIdKey)
	return ret, err
}

func (m *defaultAsTempletFormModelModel) FindOne(ctx context.Context, id string) (*AsTempletFormModel, error) {
	assetAsTempletFormModelIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormModelIdPrefix, id)
	var resp AsTempletFormModel
	err := m.QueryRowCtx(ctx, &resp, assetAsTempletFormModelIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTempletFormModelRows, m.table)
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

func (m *defaultAsTempletFormModelModel) Update(ctx context.Context, data *AsTempletFormModel) error {
	assetAsTempletFormModelIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormModelIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asTempletFormModelRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.FormName, data.ModelSheet, data.IconCls, data.Id)
	}, assetAsTempletFormModelIdKey)
	return err
}

func (m *defaultAsTempletFormModelModel) Delete(ctx context.Context, id string) error {
	assetAsTempletFormModelIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormModelIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsTempletFormModelIdKey)
	return err
}

func (m *defaultAsTempletFormModelModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsTempletFormModelIdPrefix, primary)
}

func (m *defaultAsTempletFormModelModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTempletFormModelRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsTempletFormModelModel) tableName() string {
	return m.table
}