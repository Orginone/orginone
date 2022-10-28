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
	actDmnDeploymentFieldNames          = builder.RawFieldNames(&ActDmnDeployment{})
	actDmnDeploymentRows                = strings.Join(actDmnDeploymentFieldNames, ",")
	actDmnDeploymentRowsExpectAutoSet   = strings.Join(stringx.Remove(actDmnDeploymentFieldNames, "`create_time`", "`update_time`"), ",")
	actDmnDeploymentRowsWithPlaceHolder = strings.Join(stringx.Remove(actDmnDeploymentFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActDmnDeploymentIDPrefix = "cache:asset:actDmnDeployment:iD:"
)

type (
	actDmnDeploymentModel interface {
		Insert(ctx context.Context, data *ActDmnDeployment) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActDmnDeployment, error)
		Update(ctx context.Context, data *ActDmnDeployment) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActDmnDeploymentModel struct {
		sqlc.CachedConn
		table string
	}

	ActDmnDeployment struct {
		ID                 string         `db:"ID_"`
		NAME               sql.NullString `db:"NAME_"`
		CATEGORY           sql.NullString `db:"CATEGORY_"`
		DEPLOYTIME         sql.NullTime   `db:"DEPLOY_TIME_"`
		TENANTID           sql.NullString `db:"TENANT_ID_"`
		PARENTDEPLOYMENTID sql.NullString `db:"PARENT_DEPLOYMENT_ID_"`
	}
)

func newActDmnDeploymentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActDmnDeploymentModel {
	return &defaultActDmnDeploymentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_dmn_deployment`",
	}
}

func (m *defaultActDmnDeploymentModel) Insert(ctx context.Context, data *ActDmnDeployment) (sql.Result, error) {
	assetActDmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnDeploymentIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, actDmnDeploymentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.NAME, data.CATEGORY, data.DEPLOYTIME, data.TENANTID, data.PARENTDEPLOYMENTID)
	}, assetActDmnDeploymentIDKey)
	return ret, err
}

func (m *defaultActDmnDeploymentModel) FindOne(ctx context.Context, iD string) (*ActDmnDeployment, error) {
	assetActDmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnDeploymentIDPrefix, iD)
	var resp ActDmnDeployment
	err := m.QueryRowCtx(ctx, &resp, assetActDmnDeploymentIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actDmnDeploymentRows, m.table)
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

func (m *defaultActDmnDeploymentModel) Update(ctx context.Context, data *ActDmnDeployment) error {
	assetActDmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnDeploymentIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actDmnDeploymentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.NAME, data.CATEGORY, data.DEPLOYTIME, data.TENANTID, data.PARENTDEPLOYMENTID, data.ID)
	}, assetActDmnDeploymentIDKey)
	return err
}

func (m *defaultActDmnDeploymentModel) Delete(ctx context.Context, iD string) error {
	assetActDmnDeploymentIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnDeploymentIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActDmnDeploymentIDKey)
	return err
}

func (m *defaultActDmnDeploymentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActDmnDeploymentIDPrefix, primary)
}

func (m *defaultActDmnDeploymentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actDmnDeploymentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActDmnDeploymentModel) tableName() string {
	return m.table
}