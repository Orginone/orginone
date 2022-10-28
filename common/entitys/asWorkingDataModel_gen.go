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
	asWorkingDataFieldNames          = builder.RawFieldNames(&AsWorkingData{})
	asWorkingDataRows                = strings.Join(asWorkingDataFieldNames, ",")
	asWorkingDataRowsExpectAutoSet   = strings.Join(stringx.Remove(asWorkingDataFieldNames, "`create_time`", "`update_time`"), ",")
	asWorkingDataRowsWithPlaceHolder = strings.Join(stringx.Remove(asWorkingDataFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsWorkingDataIdPrefix = "cache:asset:asWorkingData:id:"
)

type (
	asWorkingDataModel interface {
		Insert(ctx context.Context, data *AsWorkingData) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsWorkingData, error)
		Update(ctx context.Context, data *AsWorkingData) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsWorkingDataModel struct {
		sqlc.CachedConn
		table string
	}

	AsWorkingData struct {
		Id         int64         `db:"id"`          // 主码
		AppId      sql.NullInt64 `db:"app_id"`      // 应用id
		UserId     sql.NullInt64 `db:"user_id"`     // 用户id
		Tp         sql.NullInt64 `db:"tp"`          // 在途业务类型(1-用户，2-部门，3-岗位)
		CreateTime sql.NullTime  `db:"create_time"` // 创建时间
		UpdateTime sql.NullTime  `db:"update_time"` // 更新时间
		CreateUser sql.NullInt64 `db:"create_user"` // 创建用户
		UpdateUser sql.NullInt64 `db:"update_user"` // 更新用户
		IsDeleted  int64         `db:"is_deleted"`  // 是否被删除
		Status     sql.NullInt64 `db:"status"`      // 状态
	}
)

func newAsWorkingDataModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsWorkingDataModel {
	return &defaultAsWorkingDataModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_working_data`",
	}
}

func (m *defaultAsWorkingDataModel) Insert(ctx context.Context, data *AsWorkingData) (sql.Result, error) {
	assetAsWorkingDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsWorkingDataIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, asWorkingDataRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.AppId, data.UserId, data.Tp, data.CreateUser, data.UpdateUser, data.IsDeleted, data.Status)
	}, assetAsWorkingDataIdKey)
	return ret, err
}

func (m *defaultAsWorkingDataModel) FindOne(ctx context.Context, id int64) (*AsWorkingData, error) {
	assetAsWorkingDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsWorkingDataIdPrefix, id)
	var resp AsWorkingData
	err := m.QueryRowCtx(ctx, &resp, assetAsWorkingDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asWorkingDataRows, m.table)
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

func (m *defaultAsWorkingDataModel) Update(ctx context.Context, data *AsWorkingData) error {
	assetAsWorkingDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsWorkingDataIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asWorkingDataRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AppId, data.UserId, data.Tp, data.CreateUser, data.UpdateUser, data.IsDeleted, data.Status, data.Id)
	}, assetAsWorkingDataIdKey)
	return err
}

func (m *defaultAsWorkingDataModel) Delete(ctx context.Context, id int64) error {
	assetAsWorkingDataIdKey := fmt.Sprintf("%s%v", cacheAssetAsWorkingDataIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsWorkingDataIdKey)
	return err
}

func (m *defaultAsWorkingDataModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsWorkingDataIdPrefix, primary)
}

func (m *defaultAsWorkingDataModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asWorkingDataRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsWorkingDataModel) tableName() string {
	return m.table
}