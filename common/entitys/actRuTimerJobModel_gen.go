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
	actRuTimerJobFieldNames          = builder.RawFieldNames(&ActRuTimerJob{})
	actRuTimerJobRows                = strings.Join(actRuTimerJobFieldNames, ",")
	actRuTimerJobRowsExpectAutoSet   = strings.Join(stringx.Remove(actRuTimerJobFieldNames, "`create_time`", "`update_time`"), ",")
	actRuTimerJobRowsWithPlaceHolder = strings.Join(stringx.Remove(actRuTimerJobFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActRuTimerJobIDPrefix = "cache:asset:actRuTimerJob:iD:"
)

type (
	actRuTimerJobModel interface {
		Insert(ctx context.Context, data *ActRuTimerJob) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActRuTimerJob, error)
		Update(ctx context.Context, data *ActRuTimerJob) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActRuTimerJobModel struct {
		sqlc.CachedConn
		table string
	}

	ActRuTimerJob struct {
		ID                string         `db:"ID_"`
		REV               sql.NullInt64  `db:"REV_"`
		TYPE              string         `db:"TYPE_"`
		LOCKEXPTIME       sql.NullTime   `db:"LOCK_EXP_TIME_"`
		LOCKOWNER         sql.NullString `db:"LOCK_OWNER_"`
		EXCLUSIVE         sql.NullInt64  `db:"EXCLUSIVE_"`
		EXECUTIONID       sql.NullString `db:"EXECUTION_ID_"`
		PROCESSINSTANCEID sql.NullString `db:"PROCESS_INSTANCE_ID_"`
		PROCDEFID         sql.NullString `db:"PROC_DEF_ID_"`
		SCOPEID           sql.NullString `db:"SCOPE_ID_"`
		SUBSCOPEID        sql.NullString `db:"SUB_SCOPE_ID_"`
		SCOPETYPE         sql.NullString `db:"SCOPE_TYPE_"`
		SCOPEDEFINITIONID sql.NullString `db:"SCOPE_DEFINITION_ID_"`
		RETRIES           sql.NullInt64  `db:"RETRIES_"`
		EXCEPTIONSTACKID  sql.NullString `db:"EXCEPTION_STACK_ID_"`
		EXCEPTIONMSG      sql.NullString `db:"EXCEPTION_MSG_"`
		DUEDATE           sql.NullTime   `db:"DUEDATE_"`
		REPEAT            sql.NullString `db:"REPEAT_"`
		HANDLERTYPE       sql.NullString `db:"HANDLER_TYPE_"`
		HANDLERCFG        sql.NullString `db:"HANDLER_CFG_"`
		CUSTOMVALUESID    sql.NullString `db:"CUSTOM_VALUES_ID_"`
		CREATETIME        sql.NullTime   `db:"CREATE_TIME_"`
		TENANTID          string         `db:"TENANT_ID_"`
		ELEMENTID         sql.NullString `db:"ELEMENT_ID_"`
		ELEMENTNAME       sql.NullString `db:"ELEMENT_NAME_"`
		CATEGORY          sql.NullString `db:"CATEGORY_"`
		CORRELATIONID     sql.NullString `db:"CORRELATION_ID_"`
	}
)

func newActRuTimerJobModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActRuTimerJobModel {
	return &defaultActRuTimerJobModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_ru_timer_job`",
	}
}

func (m *defaultActRuTimerJobModel) Insert(ctx context.Context, data *ActRuTimerJob) (sql.Result, error) {
	assetActRuTimerJobIDKey := fmt.Sprintf("%s%v", cacheAssetActRuTimerJobIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actRuTimerJobRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.TYPE, data.LOCKEXPTIME, data.LOCKOWNER, data.EXCLUSIVE, data.EXECUTIONID, data.PROCESSINSTANCEID, data.PROCDEFID, data.SCOPEID, data.SUBSCOPEID, data.SCOPETYPE, data.SCOPEDEFINITIONID, data.RETRIES, data.EXCEPTIONSTACKID, data.EXCEPTIONMSG, data.DUEDATE, data.REPEAT, data.HANDLERTYPE, data.HANDLERCFG, data.CUSTOMVALUESID, data.CREATETIME, data.TENANTID, data.ELEMENTID, data.ELEMENTNAME, data.CATEGORY, data.CORRELATIONID)
	}, assetActRuTimerJobIDKey)
	return ret, err
}

func (m *defaultActRuTimerJobModel) FindOne(ctx context.Context, iD string) (*ActRuTimerJob, error) {
	assetActRuTimerJobIDKey := fmt.Sprintf("%s%v", cacheAssetActRuTimerJobIDPrefix, iD)
	var resp ActRuTimerJob
	err := m.QueryRowCtx(ctx, &resp, assetActRuTimerJobIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actRuTimerJobRows, m.table)
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

func (m *defaultActRuTimerJobModel) Update(ctx context.Context, data *ActRuTimerJob) error {
	assetActRuTimerJobIDKey := fmt.Sprintf("%s%v", cacheAssetActRuTimerJobIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actRuTimerJobRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.TYPE, data.LOCKEXPTIME, data.LOCKOWNER, data.EXCLUSIVE, data.EXECUTIONID, data.PROCESSINSTANCEID, data.PROCDEFID, data.SCOPEID, data.SUBSCOPEID, data.SCOPETYPE, data.SCOPEDEFINITIONID, data.RETRIES, data.EXCEPTIONSTACKID, data.EXCEPTIONMSG, data.DUEDATE, data.REPEAT, data.HANDLERTYPE, data.HANDLERCFG, data.CUSTOMVALUESID, data.CREATETIME, data.TENANTID, data.ELEMENTID, data.ELEMENTNAME, data.CATEGORY, data.CORRELATIONID, data.ID)
	}, assetActRuTimerJobIDKey)
	return err
}

func (m *defaultActRuTimerJobModel) Delete(ctx context.Context, iD string) error {
	assetActRuTimerJobIDKey := fmt.Sprintf("%s%v", cacheAssetActRuTimerJobIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActRuTimerJobIDKey)
	return err
}

func (m *defaultActRuTimerJobModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActRuTimerJobIDPrefix, primary)
}

func (m *defaultActRuTimerJobModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actRuTimerJobRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActRuTimerJobModel) tableName() string {
	return m.table
}
