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
	bladeDictFieldNames          = builder.RawFieldNames(&BladeDict{})
	bladeDictRows                = strings.Join(bladeDictFieldNames, ",")
	bladeDictRowsExpectAutoSet   = strings.Join(stringx.Remove(bladeDictFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bladeDictRowsWithPlaceHolder = strings.Join(stringx.Remove(bladeDictFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBladeDictIdPrefix = "cache:asset:bladeDict:id:"
)

type (
	bladeDictModel interface {
		Insert(ctx context.Context, data *BladeDict) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BladeDict, error)
		Update(ctx context.Context, data *BladeDict) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBladeDictModel struct {
		sqlc.CachedConn
		table string
	}

	BladeDict struct {
		Id        int64          `db:"id"`         // 主键
		ParentId  int64          `db:"parent_id"`  // 父主键
		Code      sql.NullString `db:"code"`       // 字典码
		DictKey   sql.NullInt64  `db:"dict_key"`   // 字典值
		DictValue sql.NullString `db:"dict_value"` // 字典名称
		Sort      sql.NullInt64  `db:"sort"`       // 排序
		Remark    sql.NullString `db:"remark"`     // 字典备注
		IsDeleted int64          `db:"is_deleted"` // 是否已删除
	}
)

func newBladeDictModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBladeDictModel {
	return &defaultBladeDictModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`blade_dict`",
	}
}

func (m *defaultBladeDictModel) Insert(ctx context.Context, data *BladeDict) (sql.Result, error) {
	assetBladeDictIdKey := fmt.Sprintf("%s%v", cacheAssetBladeDictIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, bladeDictRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Code, data.DictKey, data.DictValue, data.Sort, data.Remark, data.IsDeleted)
	}, assetBladeDictIdKey)
	return ret, err
}

func (m *defaultBladeDictModel) FindOne(ctx context.Context, id int64) (*BladeDict, error) {
	assetBladeDictIdKey := fmt.Sprintf("%s%v", cacheAssetBladeDictIdPrefix, id)
	var resp BladeDict
	err := m.QueryRowCtx(ctx, &resp, assetBladeDictIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeDictRows, m.table)
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

func (m *defaultBladeDictModel) Update(ctx context.Context, data *BladeDict) error {
	assetBladeDictIdKey := fmt.Sprintf("%s%v", cacheAssetBladeDictIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bladeDictRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Code, data.DictKey, data.DictValue, data.Sort, data.Remark, data.IsDeleted, data.Id)
	}, assetBladeDictIdKey)
	return err
}

func (m *defaultBladeDictModel) Delete(ctx context.Context, id int64) error {
	assetBladeDictIdKey := fmt.Sprintf("%s%v", cacheAssetBladeDictIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetBladeDictIdKey)
	return err
}

func (m *defaultBladeDictModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBladeDictIdPrefix, primary)
}

func (m *defaultBladeDictModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeDictRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBladeDictModel) tableName() string {
	return m.table
}
