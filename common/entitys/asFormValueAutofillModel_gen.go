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
	asFormValueAutofillFieldNames          = builder.RawFieldNames(&AsFormValueAutofill{})
	asFormValueAutofillRows                = strings.Join(asFormValueAutofillFieldNames, ",")
	asFormValueAutofillRowsExpectAutoSet   = strings.Join(stringx.Remove(asFormValueAutofillFieldNames, "`create_time`", "`update_time`"), ",")
	asFormValueAutofillRowsWithPlaceHolder = strings.Join(stringx.Remove(asFormValueAutofillFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsFormValueAutofillIdPrefix = "cache:asset:asFormValueAutofill:id:"
)

type (
	asFormValueAutofillModel interface {
		Insert(ctx context.Context, data *AsFormValueAutofill) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsFormValueAutofill, error)
		Update(ctx context.Context, data *AsFormValueAutofill) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsFormValueAutofillModel struct {
		sqlc.CachedConn
		table string
	}

	AsFormValueAutofill struct {
		Id          string `db:"id"`            // 主键
		SourceId    string `db:"source_id"`     // 填充内容的数据源id
		SourceType  int64  `db:"source_type"`   // 数据源类型
		SourceKey   string `db:"source_key"`    // 填充内容对应的key
		FormModelId string `db:"form_model_id"` // 表单模型id
		FormItemKey string `db:"form_item_key"` // 表单模型中表单项key属性
		IsDeleted   int64  `db:"is_deleted"`    // 0-没删除，1-被删除
	}
)

func newAsFormValueAutofillModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsFormValueAutofillModel {
	return &defaultAsFormValueAutofillModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_form_value_autofill`",
	}
}

func (m *defaultAsFormValueAutofillModel) Insert(ctx context.Context, data *AsFormValueAutofill) (sql.Result, error) {
	assetAsFormValueAutofillIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormValueAutofillIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, asFormValueAutofillRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.SourceId, data.SourceType, data.SourceKey, data.FormModelId, data.FormItemKey, data.IsDeleted)
	}, assetAsFormValueAutofillIdKey)
	return ret, err
}

func (m *defaultAsFormValueAutofillModel) FindOne(ctx context.Context, id string) (*AsFormValueAutofill, error) {
	assetAsFormValueAutofillIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormValueAutofillIdPrefix, id)
	var resp AsFormValueAutofill
	err := m.QueryRowCtx(ctx, &resp, assetAsFormValueAutofillIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormValueAutofillRows, m.table)
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

func (m *defaultAsFormValueAutofillModel) Update(ctx context.Context, data *AsFormValueAutofill) error {
	assetAsFormValueAutofillIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormValueAutofillIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asFormValueAutofillRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.SourceId, data.SourceType, data.SourceKey, data.FormModelId, data.FormItemKey, data.IsDeleted, data.Id)
	}, assetAsFormValueAutofillIdKey)
	return err
}

func (m *defaultAsFormValueAutofillModel) Delete(ctx context.Context, id string) error {
	assetAsFormValueAutofillIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormValueAutofillIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsFormValueAutofillIdKey)
	return err
}

func (m *defaultAsFormValueAutofillModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsFormValueAutofillIdPrefix, primary)
}

func (m *defaultAsFormValueAutofillModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormValueAutofillRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsFormValueAutofillModel) tableName() string {
	return m.table
}
