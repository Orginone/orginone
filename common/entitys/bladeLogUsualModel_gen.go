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
	bladeLogUsualFieldNames          = builder.RawFieldNames(&BladeLogUsual{})
	bladeLogUsualRows                = strings.Join(bladeLogUsualFieldNames, ",")
	bladeLogUsualRowsExpectAutoSet   = strings.Join(stringx.Remove(bladeLogUsualFieldNames, "`create_time`", "`update_time`"), ",")
	bladeLogUsualRowsWithPlaceHolder = strings.Join(stringx.Remove(bladeLogUsualFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBladeLogUsualIdPrefix = "cache:asset:bladeLogUsual:id:"
)

type (
	bladeLogUsualModel interface {
		Insert(ctx context.Context, data *BladeLogUsual) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BladeLogUsual, error)
		Update(ctx context.Context, data *BladeLogUsual) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBladeLogUsualModel struct {
		sqlc.CachedConn
		table string
	}

	BladeLogUsual struct {
		Id          int64          `db:"id"`           // 编号
		TenantId    string         `db:"tenant_id"`    // 租户ID
		ServiceId   sql.NullString `db:"service_id"`   // 服务ID
		ServerHost  sql.NullString `db:"server_host"`  // 服务器名
		ServerIp    sql.NullString `db:"server_ip"`    // 服务器IP地址
		Env         sql.NullString `db:"env"`          // 系统环境
		LogLevel    sql.NullString `db:"log_level"`    // 日志级别
		LogId       sql.NullString `db:"log_id"`       // 日志业务id
		LogData     sql.NullString `db:"log_data"`     // 日志数据
		Method      sql.NullString `db:"method"`       // 操作方式
		RequestUri  sql.NullString `db:"request_uri"`  // 请求URI
		RemoteIp    sql.NullString `db:"remote_ip"`    // 操作IP地址
		MethodClass sql.NullString `db:"method_class"` // 方法类
		MethodName  sql.NullString `db:"method_name"`  // 方法名
		UserAgent   sql.NullString `db:"user_agent"`   // 用户代理
		Params      sql.NullString `db:"params"`       // 操作提交的数据
		Time        sql.NullTime   `db:"time"`         // 执行时间
		CreateBy    sql.NullString `db:"create_by"`    // 创建者
		CreateTime  time.Time      `db:"create_time"`  // 创建时间
	}
)

func newBladeLogUsualModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBladeLogUsualModel {
	return &defaultBladeLogUsualModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`blade_log_usual`",
	}
}

func (m *defaultBladeLogUsualModel) Insert(ctx context.Context, data *BladeLogUsual) (sql.Result, error) {
	assetBladeLogUsualIdKey := fmt.Sprintf("%s%v", cacheAssetBladeLogUsualIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, bladeLogUsualRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.TenantId, data.ServiceId, data.ServerHost, data.ServerIp, data.Env, data.LogLevel, data.LogId, data.LogData, data.Method, data.RequestUri, data.RemoteIp, data.MethodClass, data.MethodName, data.UserAgent, data.Params, data.Time, data.CreateBy)
	}, assetBladeLogUsualIdKey)
	return ret, err
}

func (m *defaultBladeLogUsualModel) FindOne(ctx context.Context, id int64) (*BladeLogUsual, error) {
	assetBladeLogUsualIdKey := fmt.Sprintf("%s%v", cacheAssetBladeLogUsualIdPrefix, id)
	var resp BladeLogUsual
	err := m.QueryRowCtx(ctx, &resp, assetBladeLogUsualIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeLogUsualRows, m.table)
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

func (m *defaultBladeLogUsualModel) Update(ctx context.Context, data *BladeLogUsual) error {
	assetBladeLogUsualIdKey := fmt.Sprintf("%s%v", cacheAssetBladeLogUsualIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bladeLogUsualRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TenantId, data.ServiceId, data.ServerHost, data.ServerIp, data.Env, data.LogLevel, data.LogId, data.LogData, data.Method, data.RequestUri, data.RemoteIp, data.MethodClass, data.MethodName, data.UserAgent, data.Params, data.Time, data.CreateBy, data.Id)
	}, assetBladeLogUsualIdKey)
	return err
}

func (m *defaultBladeLogUsualModel) Delete(ctx context.Context, id int64) error {
	assetBladeLogUsualIdKey := fmt.Sprintf("%s%v", cacheAssetBladeLogUsualIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetBladeLogUsualIdKey)
	return err
}

func (m *defaultBladeLogUsualModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBladeLogUsualIdPrefix, primary)
}

func (m *defaultBladeLogUsualModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeLogUsualRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBladeLogUsualModel) tableName() string {
	return m.table
}
