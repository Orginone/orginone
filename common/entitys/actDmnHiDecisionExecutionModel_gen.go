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
	actDmnHiDecisionExecutionFieldNames          = builder.RawFieldNames(&ActDmnHiDecisionExecution{})
	actDmnHiDecisionExecutionRows                = strings.Join(actDmnHiDecisionExecutionFieldNames, ",")
	actDmnHiDecisionExecutionRowsExpectAutoSet   = strings.Join(stringx.Remove(actDmnHiDecisionExecutionFieldNames, "`create_time`", "`update_time`"), ",")
	actDmnHiDecisionExecutionRowsWithPlaceHolder = strings.Join(stringx.Remove(actDmnHiDecisionExecutionFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActDmnHiDecisionExecutionIDPrefix = "cache:asset:actDmnHiDecisionExecution:iD:"
)

type (
	actDmnHiDecisionExecutionModel interface {
		Insert(ctx context.Context, data *ActDmnHiDecisionExecution) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActDmnHiDecisionExecution, error)
		Update(ctx context.Context, data *ActDmnHiDecisionExecution) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActDmnHiDecisionExecutionModel struct {
		sqlc.CachedConn
		table string
	}

	ActDmnHiDecisionExecution struct {
		ID                   string         `db:"ID_"`
		DECISIONDEFINITIONID sql.NullString `db:"DECISION_DEFINITION_ID_"`
		DEPLOYMENTID         sql.NullString `db:"DEPLOYMENT_ID_"`
		STARTTIME            sql.NullTime   `db:"START_TIME_"`
		ENDTIME              sql.NullTime   `db:"END_TIME_"`
		INSTANCEID           sql.NullString `db:"INSTANCE_ID_"`
		EXECUTIONID          sql.NullString `db:"EXECUTION_ID_"`
		ACTIVITYID           sql.NullString `db:"ACTIVITY_ID_"`
		FAILED               byte           `db:"FAILED_"`
		TENANTID             sql.NullString `db:"TENANT_ID_"`
		EXECUTIONJSON        sql.NullString `db:"EXECUTION_JSON_"`
		SCOPETYPE            sql.NullString `db:"SCOPE_TYPE_"`
	}
)

func newActDmnHiDecisionExecutionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActDmnHiDecisionExecutionModel {
	return &defaultActDmnHiDecisionExecutionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_dmn_hi_decision_execution`",
	}
}

func (m *defaultActDmnHiDecisionExecutionModel) Insert(ctx context.Context, data *ActDmnHiDecisionExecution) (sql.Result, error) {
	assetActDmnHiDecisionExecutionIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnHiDecisionExecutionIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actDmnHiDecisionExecutionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.DECISIONDEFINITIONID, data.DEPLOYMENTID, data.STARTTIME, data.ENDTIME, data.INSTANCEID, data.EXECUTIONID, data.ACTIVITYID, data.FAILED, data.TENANTID, data.EXECUTIONJSON, data.SCOPETYPE)
	}, assetActDmnHiDecisionExecutionIDKey)
	return ret, err
}

func (m *defaultActDmnHiDecisionExecutionModel) FindOne(ctx context.Context, iD string) (*ActDmnHiDecisionExecution, error) {
	assetActDmnHiDecisionExecutionIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnHiDecisionExecutionIDPrefix, iD)
	var resp ActDmnHiDecisionExecution
	err := m.QueryRowCtx(ctx, &resp, assetActDmnHiDecisionExecutionIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actDmnHiDecisionExecutionRows, m.table)
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

func (m *defaultActDmnHiDecisionExecutionModel) Update(ctx context.Context, data *ActDmnHiDecisionExecution) error {
	assetActDmnHiDecisionExecutionIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnHiDecisionExecutionIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actDmnHiDecisionExecutionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DECISIONDEFINITIONID, data.DEPLOYMENTID, data.STARTTIME, data.ENDTIME, data.INSTANCEID, data.EXECUTIONID, data.ACTIVITYID, data.FAILED, data.TENANTID, data.EXECUTIONJSON, data.SCOPETYPE, data.ID)
	}, assetActDmnHiDecisionExecutionIDKey)
	return err
}

func (m *defaultActDmnHiDecisionExecutionModel) Delete(ctx context.Context, iD string) error {
	assetActDmnHiDecisionExecutionIDKey := fmt.Sprintf("%s%v", cacheAssetActDmnHiDecisionExecutionIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActDmnHiDecisionExecutionIDKey)
	return err
}

func (m *defaultActDmnHiDecisionExecutionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActDmnHiDecisionExecutionIDPrefix, primary)
}

func (m *defaultActDmnHiDecisionExecutionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actDmnHiDecisionExecutionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActDmnHiDecisionExecutionModel) tableName() string {
	return m.table
}
