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
	asRedeployDataFieldNames          = builder.RawFieldNames(&AsRedeployData{})
	asRedeployDataRows                = strings.Join(asRedeployDataFieldNames, ",")
	asRedeployDataRowsExpectAutoSet   = strings.Join(stringx.Remove(asRedeployDataFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asRedeployDataRowsWithPlaceHolder = strings.Join(stringx.Remove(asRedeployDataFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsRedeployDataIdPrefix = "cache:asset:asRedeployData:id:"
)

type (
	asRedeployDataModel interface {
		Insert(ctx context.Context, data *AsRedeployData) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsRedeployData, error)
		Update(ctx context.Context, data *AsRedeployData) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsRedeployDataModel struct {
		sqlc.CachedConn
		table string
	}

	AsRedeployData struct {
		Id          int64          `db:"id"`          // 主键
		AppId       sql.NullInt64  `db:"app_id"`      // 应用id
		Application sql.NullString `db:"application"` // 应用详情
		CreateTime  time.Time      `db:"create_time"` // 创建时间
		UpdateTime  time.Time      `db:"update_time"` // 更新时间
		CreateUser  sql.NullInt64  `db:"create_user"` // 创建用户
		UpdateUser  sql.NullInt64  `db:"update_user"` // 更新用户
		IsDeleted   int64          `db:"is_deleted"`  // 是否被删除
		Status      int64          `db:"status"`      // 状态(0-申请，1-通过，2-拒绝)
	}
)

func newAsRedeployDataModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsRedeployDataModel {
	return &defaultAsRedeployDataModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_redeploy_data`",
	}
}

func (m *defaultAsRedeployDataModel) Insert(ctx context.Context, data *AsRedeployData) (sql.Result, error) {
	assetAsRedeployDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsRedeployDataIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, asRedeployDataRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AppId, data.Application, data.CreateUser, data.UpdateUser, data.IsDeleted, data.Status)
	}, assetAsRedeployDataIdKey)
	return ret, err
}

func (m *defaultAsRedeployDataModel) FindOne(ctx context.Context, id int64) (*AsRedeployData, error) {
	assetAsRedeployDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsRedeployDataIdPrefix, id)
	var resp AsRedeployData
	err := m.QueryRowCtx(ctx, &resp, assetAsRedeployDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asRedeployDataRows, m.table)
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

func (m *defaultAsRedeployDataModel) Update(ctx context.Context, data *AsRedeployData) error {
	assetAsRedeployDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsRedeployDataIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asRedeployDataRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AppId, data.Application, data.CreateUser, data.UpdateUser, data.IsDeleted, data.Status, data.Id)
	}, assetAsRedeployDataIdKey)
	return err
}

func (m *defaultAsRedeployDataModel) Delete(ctx context.Context, id int64) error {
	assetAsRedeployDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsRedeployDataIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsRedeployDataIdKey)
	return err
}

func (m *defaultAsRedeployDataModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsRedeployDataIdPrefix, primary)
}

func (m *defaultAsRedeployDataModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asRedeployDataRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsRedeployDataModel) tableName() string {
	return m.table
}
