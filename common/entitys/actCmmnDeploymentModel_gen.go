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
	actCmmnDeploymentFieldNames          = builder.RawFieldNames(&ActCmmnDeployment{})
	actCmmnDeploymentRows                = strings.Join(actCmmnDeploymentFieldNames, ",")
	actCmmnDeploymentRowsExpectAutoSet   = strings.Join(stringx.Remove(actCmmnDeploymentFieldNames, "`create_time`", "`update_time`"), ",")
	actCmmnDeploymentRowsWithPlaceHolder = strings.Join(stringx.Remove(actCmmnDeploymentFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActCmmnDeploymentIDPrefix = "cache:asset:actCmmnDeployment:iD:"
)

type (
	actCmmnDeploymentModel interface {
		Insert(ctx context.Context, data *ActCmmnDeployment) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActCmmnDeployment, error)
		Update(ctx context.Context, data *ActCmmnDeployment) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActCmmnDeploymentModel struct {
		sqlc.CachedConn
		table string
	}

	ActCmmnDeployment struct {
		ID                 string         `db:"ID_"`
		NAME               sql.NullString `db:"NAME_"`
		CATEGORY           sql.NullString `db:"CATEGORY_"`
		KEY                sql.NullString `db:"KEY_"`
		DEPLOYTIME         sql.NullTime   `db:"DEPLOY_TIME_"`
		PARENTDEPLOYMENTID sql.NullString `db:"PARENT_DEPLOYMENT_ID_"`
		TENANTID           string         `db:"TENANT_ID_"`
	}
)

func newActCmmnDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActCmmnDeploymentModel {
	return &defaultActCmmnDeploymentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_cmmn_deployment`",
	}
}

func (m *defaultActCmmnDeploymentModel) Insert(ctx context.Context, data *ActCmmnDeployment) (sql.Result, error) {
	assetActCmmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnDeploymentIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, actCmmnDeploymentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.NAME, data.CATEGORY, data.KEY, data.DEPLOYTIME, data.PARENTDEPLOYMENTID, data.TENANTID)
	}, assetActCmmnDeploymentIDKey)
	return ret, err
}

func (m *defaultActCmmnDeploymentModel) FindOne(ctx context.Context, iD string) (*ActCmmnDeployment, error) {
	assetActCmmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnDeploymentIDPrefix, iD)
	var resp ActCmmnDeployment
	err := m.QueryRowCtx(ctx, &resp, assetActCmmnDeploymentIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnDeploymentRows, m.table)
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

func (m *defaultActCmmnDeploymentModel) Update(ctx context.Context, data *ActCmmnDeployment) error {
	assetActCmmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnDeploymentIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actCmmnDeploymentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.NAME, data.CATEGORY, data.KEY, data.DEPLOYTIME, data.PARENTDEPLOYMENTID, data.TENANTID, data.ID)
	}, assetActCmmnDeploymentIDKey)
	return err
}

func (m *defaultActCmmnDeploymentModel) Delete(ctx context.Context, iD string) error {
	assetActCmmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnDeploymentIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActCmmnDeploymentIDKey)
	return err
}

func (m *defaultActCmmnDeploymentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActCmmnDeploymentIDPrefix, primary)
}

func (m *defaultActCmmnDeploymentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnDeploymentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActCmmnDeploymentModel) tableName() string {
	return m.table
}
