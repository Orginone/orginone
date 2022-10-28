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
	actReModelFieldNames          = builder.RawFieldNames(&ActReModel{})
	actReModelRows                = strings.Join(actReModelFieldNames, ",")
	actReModelRowsExpectAutoSet   = strings.Join(stringx.Remove(actReModelFieldNames, "`create_time`", "`update_time`"), ",")
	actReModelRowsWithPlaceHolder = strings.Join(stringx.Remove(actReModelFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActReModelIDPrefix = "cache:asset:actReModel:iD:"
)

type (
	actReModelModel interface {
		Insert(ctx context.Context, data *ActReModel) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActReModel, error)
		Update(ctx context.Context, data *ActReModel) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActReModelModel struct {
		sqlc.CachedConn
		table string
	}

	ActReModel struct {
		ID                       string         `db:"ID_"`
		REV                      sql.NullInt64  `db:"REV_"`
		NAME                     sql.NullString `db:"NAME_"`
		KEY                      sql.NullString `db:"KEY_"`
		CATEGORY                 sql.NullString `db:"CATEGORY_"`
		CREATETIME               sql.NullTime   `db:"CREATE_TIME_"`
		LASTUPDATETIME           sql.NullTime   `db:"LAST_UPDATE_TIME_"`
		VERSION                  sql.NullInt64  `db:"VERSION_"`
		METAINFO                 sql.NullString `db:"META_INFO_"`
		DEPLOYMENTID             sql.NullString `db:"DEPLOYMENT_ID_"`
		EDITORSOURCEVALUEID      sql.NullString `db:"EDITOR_SOURCE_VALUE_ID_"`
		EDITORSOURCEEXTRAVALUEID sql.NullString `db:"EDITOR_SOURCE_EXTRA_VALUE_ID_"`
		TENANTID                 string         `db:"TENANT_ID_"`
	}
)

func newActReModelModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActReModelModel {
	return &defaultActReModelModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_re_model`",
	}
}

func (m *defaultActReModelModel) Insert(ctx context.Context, data *ActReModel) (sql.Result, error) {
	assetActReModelIDKey := fmt.Sprintf("%s%v", cacheAssetActReModelIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actReModelRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.NAME, data.KEY, data.CATEGORY, data.CREATETIME, data.LASTUPDATETIME, data.VERSION, data.METAINFO, data.DEPLOYMENTID, data.EDITORSOURCEVALUEID, data.EDITORSOURCEEXTRAVALUEID, data.TENANTID)
	}, assetActReModelIDKey)
	return ret, err
}

func (m *defaultActReModelModel) FindOne(ctx context.Context, iD string) (*ActReModel, error) {
	assetActReModelIDKey := fmt.Sprintf("%s%v", cacheAssetActReModelIDPrefix, iD)
	var resp ActReModel
	err := m.QueryRowCtx(ctx, &resp, assetActReModelIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actReModelRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, iD)
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

func (m *defaultActReModelModel) Update(ctx context.Context, data *ActReModel) error {
	assetActReModelIDKey := fmt.Sprintf("%s%v", cacheAssetActReModelIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actReModelRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.NAME, data.KEY, data.CATEGORY, data.CREATETIME, data.LASTUPDATETIME, data.VERSION, data.METAINFO, data.DEPLOYMENTID, data.EDITORSOURCEVALUEID, data.EDITORSOURCEEXTRAVALUEID, data.TENANTID, data.ID)
	}, assetActReModelIDKey)
	return err
}

func (m *defaultActReModelModel) Delete(ctx context.Context, iD string) error {
	assetActReModelIDKey := fmt.Sprintf("%s%v", cacheAssetActReModelIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActReModelIDKey)
	return err
}

func (m *defaultActReModelModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActReModelIDPrefix, primary)
}

func (m *defaultActReModelModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actReModelRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActReModelModel) tableName() string {
	return m.table
}
