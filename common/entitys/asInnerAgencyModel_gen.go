// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	asInnerAgencyFieldNames          = builder.RawFieldNames(&AsInnerAgency{})
	asInnerAgencyRows                = strings.Join(asInnerAgencyFieldNames, ",")
	asInnerAgencyRowsExpectAutoSet   = strings.Join(stringx.Remove(asInnerAgencyFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asInnerAgencyRowsWithPlaceHolder = strings.Join(stringx.Remove(asInnerAgencyFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsInnerAgencyIdPrefix = "cache:asset:asInnerAgency:id:"
)

type (
	asInnerAgencyModel interface {
		Insert(ctx context.Context, data *AsInnerAgency) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsInnerAgency, error)
		Update(ctx context.Context, data *AsInnerAgency) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsInnerAgencyModel struct {
		sqlc.CachedConn
		table string
	}

	AsInnerAgency struct {
		Id         int64          `db:"id"`
		AgencyName string         `db:"agency_name"` // 机构名称
		AgencyCode sql.NullString `db:"agency_code"` // 部门编码(系统生成）
		TenantCode string         `db:"tenant_code"` // 所属租户ID
		ParentId   int64          `db:"parent_id"`   // 父Id
		IsDeleted  int64          `db:"is_deleted"`  // 是否被删除；0-有效；1-删除；
		UpdateUser sql.NullInt64  `db:"update_user"` // 更新人
		UpdateTime time.Time      `db:"update_time"` // 更新时间
		CreateTime time.Time      `db:"create_time"` // 创建时间
		CreateUser sql.NullInt64  `db:"create_user"` // 创建人
		Status     int64          `db:"status"`      // 状态
	}
)

func newAsInnerAgencyModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsInnerAgencyModel {
	return &defaultAsInnerAgencyModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_inner_agency`",
	}
}

func (m *defaultAsInnerAgencyModel) Insert(ctx context.Context, data *AsInnerAgency) (sql.Result, error) {
	assetAsInnerAgencyIdKey := fmt.Sprintf("%s%v", cacheAssetAsInnerAgencyIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, asInnerAgencyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AgencyName, data.AgencyCode, data.TenantCode, data.ParentId, data.IsDeleted, data.UpdateUser, data.CreateUser, data.Status)
	}, assetAsInnerAgencyIdKey)
	return ret, err
}

func (m *defaultAsInnerAgencyModel) FindOne(ctx context.Context, id int64) (*AsInnerAgency, error) {
	assetAsInnerAgencyIdKey := fmt.Sprintf("%s%v", cacheAssetAsInnerAgencyIdPrefix, id)
	var resp AsInnerAgency
	err := m.QueryRowCtx(ctx, &resp, assetAsInnerAgencyIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asInnerAgencyRows, m.table)
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

func (m *defaultAsInnerAgencyModel) Update(ctx context.Context, data *AsInnerAgency) error {
	assetAsInnerAgencyIdKey := fmt.Sprintf("%s%v", cacheAssetAsInnerAgencyIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asInnerAgencyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AgencyName, data.AgencyCode, data.TenantCode, data.ParentId, data.IsDeleted, data.UpdateUser, data.CreateUser, data.Status, data.Id)
	}, assetAsInnerAgencyIdKey)
	return err
}

func (m *defaultAsInnerAgencyModel) Delete(ctx context.Context, id int64) error {
	assetAsInnerAgencyIdKey := fmt.Sprintf("%s%v", cacheAssetAsInnerAgencyIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsInnerAgencyIdKey)
	return err
}

func (m *defaultAsInnerAgencyModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsInnerAgencyIdPrefix, primary)
}

func (m *defaultAsInnerAgencyModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asInnerAgencyRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsInnerAgencyModel) tableName() string {
	return m.table
}
