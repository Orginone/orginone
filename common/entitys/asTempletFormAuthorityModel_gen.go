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
	asTempletFormAuthorityFieldNames          = builder.RawFieldNames(&AsTempletFormAuthority{})
	asTempletFormAuthorityRows                = strings.Join(asTempletFormAuthorityFieldNames, ",")
	asTempletFormAuthorityRowsExpectAutoSet   = strings.Join(stringx.Remove(asTempletFormAuthorityFieldNames, "`create_time`", "`update_time`"), ",")
	asTempletFormAuthorityRowsWithPlaceHolder = strings.Join(stringx.Remove(asTempletFormAuthorityFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsTempletFormAuthorityIdPrefix = "cache:asset:asTempletFormAuthority:id:"
)

type (
	asTempletFormAuthorityModel interface {
		Insert(ctx context.Context, data *AsTempletFormAuthority) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsTempletFormAuthority, error)
		Update(ctx context.Context, data *AsTempletFormAuthority) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsTempletFormAuthorityModel struct {
		sqlc.CachedConn
		table string
	}

	AsTempletFormAuthority struct {
		Id                 string `db:"id"` // 表单项权限模板的Id
		TempletProcModelId string `db:"templet_proc_model_id"`
		ActId              string `db:"act_id"`
		FormItemKey        string `db:"form_item_key"`
		Authority          int64  `db:"authority"` // 1不可见，2可见+不可编辑，3可见+可编辑，4必填
	}
)

func newAsTempletFormAuthorityModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsTempletFormAuthorityModel {
	return &defaultAsTempletFormAuthorityModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_templet_form_authority`",
	}
}

func (m *defaultAsTempletFormAuthorityModel) Insert(ctx context.Context, data *AsTempletFormAuthority) (sql.Result, error) {
	assetAsTempletFormAuthorityIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormAuthorityIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, asTempletFormAuthorityRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.TempletProcModelId, data.ActId, data.FormItemKey, data.Authority)
	}, assetAsTempletFormAuthorityIdKey)
	return ret, err
}

func (m *defaultAsTempletFormAuthorityModel) FindOne(ctx context.Context, id string) (*AsTempletFormAuthority, error) {
	assetAsTempletFormAuthorityIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormAuthorityIdPrefix, id)
	var resp AsTempletFormAuthority
	err := m.QueryRowCtx(ctx, &resp, assetAsTempletFormAuthorityIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTempletFormAuthorityRows, m.table)
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

func (m *defaultAsTempletFormAuthorityModel) Update(ctx context.Context, data *AsTempletFormAuthority) error {
	assetAsTempletFormAuthorityIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormAuthorityIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asTempletFormAuthorityRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TempletProcModelId, data.ActId, data.FormItemKey, data.Authority, data.Id)
	}, assetAsTempletFormAuthorityIdKey)
	return err
}

func (m *defaultAsTempletFormAuthorityModel) Delete(ctx context.Context, id string) error {
	assetAsTempletFormAuthorityIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletFormAuthorityIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsTempletFormAuthorityIdKey)
	return err
}

func (m *defaultAsTempletFormAuthorityModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsTempletFormAuthorityIdPrefix, primary)
}

func (m *defaultAsTempletFormAuthorityModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTempletFormAuthorityRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsTempletFormAuthorityModel) tableName() string {
	return m.table
}
