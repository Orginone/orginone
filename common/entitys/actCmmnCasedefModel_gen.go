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
	actCmmnCasedefFieldNames          = builder.RawFieldNames(&ActCmmnCasedef{})
	actCmmnCasedefRows                = strings.Join(actCmmnCasedefFieldNames, ",")
	actCmmnCasedefRowsExpectAutoSet   = strings.Join(stringx.Remove(actCmmnCasedefFieldNames, "`create_time`", "`update_time`"), ",")
	actCmmnCasedefRowsWithPlaceHolder = strings.Join(stringx.Remove(actCmmnCasedefFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActCmmnCasedefIDPrefix = "cache:asset:actCmmnCasedef:iD:"
)

type (
	actCmmnCasedefModel interface {
		Insert(ctx context.Context, data *ActCmmnCasedef) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActCmmnCasedef, error)
		Update(ctx context.Context, data *ActCmmnCasedef) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActCmmnCasedefModel struct {
		sqlc.CachedConn
		table string
	}

	ActCmmnCasedef struct {
		ID                   string         `db:"ID_"`
		REV                  int64          `db:"REV_"`
		NAME                 sql.NullString `db:"NAME_"`
		KEY                  string         `db:"KEY_"`
		VERSION              int64          `db:"VERSION_"`
		CATEGORY             sql.NullString `db:"CATEGORY_"`
		DEPLOYMENTID         sql.NullString `db:"DEPLOYMENT_ID_"`
		RESOURCENAME         sql.NullString `db:"RESOURCE_NAME_"`
		DESCRIPTION          sql.NullString `db:"DESCRIPTION_"`
		HASGRAPHICALNOTATION byte           `db:"HAS_GRAPHICAL_NOTATION_"`
		TENANTID             string         `db:"TENANT_ID_"`
		DGRMRESOURCENAME     sql.NullString `db:"DGRM_RESOURCE_NAME_"`
		HASSTARTFORMKEY      byte           `db:"HAS_START_FORM_KEY_"`
	}
)

func newActCmmnCasedefModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActCmmnCasedefModel {
	return &defaultActCmmnCasedefModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_cmmn_casedef`",
	}
}

func (m *defaultActCmmnCasedefModel) Insert(ctx context.Context, data *ActCmmnCasedef) (sql.Result, error) {
	assetActCmmnCasedefIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnCasedefIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actCmmnCasedefRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.NAME, data.KEY, data.VERSION, data.CATEGORY, data.DEPLOYMENTID, data.RESOURCENAME, data.DESCRIPTION, data.HASGRAPHICALNOTATION, data.TENANTID, data.DGRMRESOURCENAME, data.HASSTARTFORMKEY)
	}, assetActCmmnCasedefIDKey)
	return ret, err
}

func (m *defaultActCmmnCasedefModel) FindOne(ctx context.Context, iD string) (*ActCmmnCasedef, error) {
	assetActCmmnCasedefIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnCasedefIDPrefix, iD)
	var resp ActCmmnCasedef
	err := m.QueryRowCtx(ctx, &resp, assetActCmmnCasedefIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnCasedefRows, m.table)
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

func (m *defaultActCmmnCasedefModel) Update(ctx context.Context, data *ActCmmnCasedef) error {
	assetActCmmnCasedefIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnCasedefIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actCmmnCasedefRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.NAME, data.KEY, data.VERSION, data.CATEGORY, data.DEPLOYMENTID, data.RESOURCENAME, data.DESCRIPTION, data.HASGRAPHICALNOTATION, data.TENANTID, data.DGRMRESOURCENAME, data.HASSTARTFORMKEY, data.ID)
	}, assetActCmmnCasedefIDKey)
	return err
}

func (m *defaultActCmmnCasedefModel) Delete(ctx context.Context, iD string) error {
	assetActCmmnCasedefIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnCasedefIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActCmmnCasedefIDKey)
	return err
}

func (m *defaultActCmmnCasedefModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActCmmnCasedefIDPrefix, primary)
}

func (m *defaultActCmmnCasedefModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnCasedefRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActCmmnCasedefModel) tableName() string {
	return m.table
}
