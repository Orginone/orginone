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
	asAppJobFieldNames          = builder.RawFieldNames(&AsAppJob{})
	asAppJobRows                = strings.Join(asAppJobFieldNames, ",")
	asAppJobRowsExpectAutoSet   = strings.Join(stringx.Remove(asAppJobFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asAppJobRowsWithPlaceHolder = strings.Join(stringx.Remove(asAppJobFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsAppJobIdPrefix = "cache:asset:asAppJob:id:"
)

type (
	asAppJobModel interface {
		Insert(ctx context.Context, data *AsAppJob) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsAppJob, error)
		Update(ctx context.Context, data *AsAppJob) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsAppJobModel struct {
		sqlc.CachedConn
		table string
	}

	AsAppJob struct {
		Id         int64         `db:"id"`          // 主键
		CreateUser sql.NullInt64 `db:"create_user"` // 创建人
		CreateTime sql.NullTime  `db:"create_time"` // 创建时间
		UpdateUser sql.NullInt64 `db:"update_user"` // 修改人
		UpdateTime sql.NullTime  `db:"update_time"` // 修改时间
		Status     sql.NullInt64 `db:"status"`      // 状态
		IsDeleted  sql.NullInt64 `db:"is_deleted"`  // 是否已删除
		AppId      sql.NullInt64 `db:"app_id"`      // 应用id
		JobId      sql.NullInt64 `db:"job_id"`      // 岗位id
	}
)

func newAsAppJobModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsAppJobModel {
	return &defaultAsAppJobModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_app_job`",
	}
}

func (m *defaultAsAppJobModel) Insert(ctx context.Context, data *AsAppJob) (sql.Result, error) {
	assetAsAppJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsAppJobIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, asAppJobRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.AppId, data.JobId)
	}, assetAsAppJobIdKey)
	return ret, err
}

func (m *defaultAsAppJobModel) FindOne(ctx context.Context, id int64) (*AsAppJob, error) {
	assetAsAppJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsAppJobIdPrefix, id)
	var resp AsAppJob
	err := m.QueryRowCtx(ctx, &resp, assetAsAppJobIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asAppJobRows, m.table)
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

func (m *defaultAsAppJobModel) Update(ctx context.Context, data *AsAppJob) error {
	assetAsAppJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsAppJobIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asAppJobRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.AppId, data.JobId, data.Id)
	}, assetAsAppJobIdKey)
	return err
}

func (m *defaultAsAppJobModel) Delete(ctx context.Context, id int64) error {
	assetAsAppJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsAppJobIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsAppJobIdKey)
	return err
}

func (m *defaultAsAppJobModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsAppJobIdPrefix, primary)
}

func (m *defaultAsAppJobModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asAppJobRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsAppJobModel) tableName() string {
	return m.table
}
