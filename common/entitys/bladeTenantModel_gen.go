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
	bladeTenantFieldNames          = builder.RawFieldNames(&BladeTenant{})
	bladeTenantRows                = strings.Join(bladeTenantFieldNames, ",")
	bladeTenantRowsExpectAutoSet   = strings.Join(stringx.Remove(bladeTenantFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bladeTenantRowsWithPlaceHolder = strings.Join(stringx.Remove(bladeTenantFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBladeTenantIdPrefix = "cache:asset:bladeTenant:id:"
)

type (
	bladeTenantModel interface {
		Insert(ctx context.Context, data *BladeTenant) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BladeTenant, error)
		Update(ctx context.Context, data *BladeTenant) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBladeTenantModel struct {
		sqlc.CachedConn
		table string
	}

	BladeTenant struct {
		Id            int64          `db:"id"`             // 主键
		TenantCode    string         `db:"tenant_code"`    // 租户编号
		TenantName    string         `db:"tenant_name"`    // 租户名称
		Linkman       sql.NullString `db:"linkman"`        // 联系人
		ContactNumber sql.NullString `db:"contact_number"` // 联系电话
		Address       sql.NullString `db:"address"`        // 联系地址
		CreateUser    sql.NullInt64  `db:"create_user"`    // 创建人
		CreateTime    sql.NullTime   `db:"create_time"`    // 创建时间
		UpdateUser    sql.NullInt64  `db:"update_user"`    // 修改人
		UpdateTime    sql.NullTime   `db:"update_time"`    // 修改时间
		Status        sql.NullInt64  `db:"status"`         // 状态
		IsDeleted     int64          `db:"is_deleted"`     // 是否已删除
	}
)

func newBladeTenantModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBladeTenantModel {
	return &defaultBladeTenantModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`blade_tenant`",
	}
}

func (m *defaultBladeTenantModel) Insert(ctx context.Context, data *BladeTenant) (sql.Result, error) {
	assetBladeTenantIdKey := fmt.Sprintf("%s%v", cacheAssetBladeTenantIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, bladeTenantRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TenantCode, data.TenantName, data.Linkman, data.ContactNumber, data.Address, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted)
	}, assetBladeTenantIdKey)
	return ret, err
}

func (m *defaultBladeTenantModel) FindOne(ctx context.Context, id int64) (*BladeTenant, error) {
	assetBladeTenantIdKey := fmt.Sprintf("%s%v", cacheAssetBladeTenantIdPrefix, id)
	var resp BladeTenant
	err := m.QueryRowCtx(ctx, &resp, assetBladeTenantIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeTenantRows, m.table)
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

func (m *defaultBladeTenantModel) Update(ctx context.Context, data *BladeTenant) error {
	assetBladeTenantIdKey := fmt.Sprintf("%s%v", cacheAssetBladeTenantIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bladeTenantRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TenantCode, data.TenantName, data.Linkman, data.ContactNumber, data.Address, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.Id)
	}, assetBladeTenantIdKey)
	return err
}

func (m *defaultBladeTenantModel) Delete(ctx context.Context, id int64) error {
	assetBladeTenantIdKey := fmt.Sprintf("%s%v", cacheAssetBladeTenantIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetBladeTenantIdKey)
	return err
}

func (m *defaultBladeTenantModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBladeTenantIdPrefix, primary)
}

func (m *defaultBladeTenantModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeTenantRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBladeTenantModel) tableName() string {
	return m.table
}
