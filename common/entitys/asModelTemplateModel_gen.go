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
	asModelTemplateFieldNames          = builder.RawFieldNames(&AsModelTemplate{})
	asModelTemplateRows                = strings.Join(asModelTemplateFieldNames, ",")
	asModelTemplateRowsExpectAutoSet   = strings.Join(stringx.Remove(asModelTemplateFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asModelTemplateRowsWithPlaceHolder = strings.Join(stringx.Remove(asModelTemplateFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsModelTemplateIdPrefix = "cache:asset:asModelTemplate:id:"
)

type (
	asModelTemplateModel interface {
		Insert(ctx context.Context, data *AsModelTemplate) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsModelTemplate, error)
		Update(ctx context.Context, data *AsModelTemplate) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsModelTemplateModel struct {
		sqlc.CachedConn
		table string
	}

	AsModelTemplate struct {
		FormId     sql.NullString `db:"form_id"`     // 表单id
		TemplateId sql.NullInt64  `db:"template_id"` // 模板id
		Id         int64          `db:"id"`
	}
)

func newAsModelTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsModelTemplateModel {
	return &defaultAsModelTemplateModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_model_template`",
	}
}

func (m *defaultAsModelTemplateModel) Insert(ctx context.Context, data *AsModelTemplate) (sql.Result, error) {
	assetAsModelTemplateIdKey := fmt.Sprintf("%s%v", cacheAssetAsModelTemplateIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, asModelTemplateRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.FormId, data.TemplateId)
	}, assetAsModelTemplateIdKey)
	return ret, err
}

func (m *defaultAsModelTemplateModel) FindOne(ctx context.Context, id int64) (*AsModelTemplate, error) {
	assetAsModelTemplateIdKey := fmt.Sprintf("%s%v", cacheAssetAsModelTemplateIdPrefix, id)
	var resp AsModelTemplate
	err := m.QueryRowCtx(ctx, &resp, assetAsModelTemplateIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asModelTemplateRows, m.table)
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

func (m *defaultAsModelTemplateModel) Update(ctx context.Context, data *AsModelTemplate) error {
	assetAsModelTemplateIdKey := fmt.Sprintf("%s%v", cacheAssetAsModelTemplateIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asModelTemplateRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.FormId, data.TemplateId, data.Id)
	}, assetAsModelTemplateIdKey)
	return err
}

func (m *defaultAsModelTemplateModel) Delete(ctx context.Context, id int64) error {
	assetAsModelTemplateIdKey := fmt.Sprintf("%s%v", cacheAssetAsModelTemplateIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsModelTemplateIdKey)
	return err
}

func (m *defaultAsModelTemplateModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsModelTemplateIdPrefix, primary)
}

func (m *defaultAsModelTemplateModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asModelTemplateRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsModelTemplateModel) tableName() string {
	return m.table
}
