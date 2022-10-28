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
	asJobPersonFieldNames          = builder.RawFieldNames(&AsJobPerson{})
	asJobPersonRows                = strings.Join(asJobPersonFieldNames, ",")
	asJobPersonRowsExpectAutoSet   = strings.Join(stringx.Remove(asJobPersonFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asJobPersonRowsWithPlaceHolder = strings.Join(stringx.Remove(asJobPersonFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsJobPersonIdPrefix = "cache:asset:asJobPerson:id:"
)

type (
	asJobPersonModel interface {
		Insert(ctx context.Context, data *AsJobPerson) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsJobPerson, error)
		Update(ctx context.Context, data *AsJobPerson) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsJobPersonModel struct {
		sqlc.CachedConn
		table string
	}

	AsJobPerson struct {
		JobId    sql.NullInt64 `db:"job_id"`
		PersonId sql.NullInt64 `db:"person_id"`
		Id       int64         `db:"id"`
	}
)

func newAsJobPersonModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsJobPersonModel {
	return &defaultAsJobPersonModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_job_person`",
	}
}

func (m *defaultAsJobPersonModel) Insert(ctx context.Context, data *AsJobPerson) (sql.Result, error) {
	assetAsJobPersonIdKey := fmt.Sprintf("%s%v", cacheAssetAsJobPersonIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, asJobPersonRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.JobId, data.PersonId)
	}, assetAsJobPersonIdKey)
	return ret, err
}

func (m *defaultAsJobPersonModel) FindOne(ctx context.Context, id int64) (*AsJobPerson, error) {
	assetAsJobPersonIdKey := fmt.Sprintf("%s%v", cacheAssetAsJobPersonIdPrefix, id)
	var resp AsJobPerson
	err := m.QueryRowCtx(ctx, &resp, assetAsJobPersonIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asJobPersonRows, m.table)
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

func (m *defaultAsJobPersonModel) Update(ctx context.Context, data *AsJobPerson) error {
	assetAsJobPersonIdKey := fmt.Sprintf("%s%v", cacheAssetAsJobPersonIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asJobPersonRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.JobId, data.PersonId, data.Id)
	}, assetAsJobPersonIdKey)
	return err
}

func (m *defaultAsJobPersonModel) Delete(ctx context.Context, id int64) error {
	assetAsJobPersonIdKey := fmt.Sprintf("%s%v", cacheAssetAsJobPersonIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsJobPersonIdKey)
	return err
}

func (m *defaultAsJobPersonModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsJobPersonIdPrefix, primary)
}

func (m *defaultAsJobPersonModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asJobPersonRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsJobPersonModel) tableName() string {
	return m.table
}
