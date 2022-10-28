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
	actRuEventSubscrFieldNames          = builder.RawFieldNames(&ActRuEventSubscr{})
	actRuEventSubscrRows                = strings.Join(actRuEventSubscrFieldNames, ",")
	actRuEventSubscrRowsExpectAutoSet   = strings.Join(stringx.Remove(actRuEventSubscrFieldNames, "`create_time`", "`update_time`"), ",")
	actRuEventSubscrRowsWithPlaceHolder = strings.Join(stringx.Remove(actRuEventSubscrFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActRuEventSubscrIDPrefix = "cache:asset:actRuEventSubscr:iD:"
)

type (
	actRuEventSubscrModel interface {
		Insert(ctx context.Context, data *ActRuEventSubscr) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActRuEventSubscr, error)
		Update(ctx context.Context, data *ActRuEventSubscr) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActRuEventSubscrModel struct {
		sqlc.CachedConn
		table string
	}

	ActRuEventSubscr struct {
		ID                string         `db:"ID_"`
		REV               sql.NullInt64  `db:"REV_"`
		EVENTTYPE         string         `db:"EVENT_TYPE_"`
		EVENTNAME         sql.NullString `db:"EVENT_NAME_"`
		EXECUTIONID       sql.NullString `db:"EXECUTION_ID_"`
		PROCINSTID        sql.NullString `db:"PROC_INST_ID_"`
		ACTIVITYID        sql.NullString `db:"ACTIVITY_ID_"`
		CONFIGURATION     sql.NullString `db:"CONFIGURATION_"`
		CREATED           time.Time      `db:"CREATED_"`
		PROCDEFID         sql.NullString `db:"PROC_DEF_ID_"`
		TENANTID          string         `db:"TENANT_ID_"`
		SUBSCOPEID        sql.NullString `db:"SUB_SCOPE_ID_"`
		SCOPEID           sql.NullString `db:"SCOPE_ID_"`
		SCOPEDEFINITIONID sql.NullString `db:"SCOPE_DEFINITION_ID_"`
		SCOPETYPE         sql.NullString `db:"SCOPE_TYPE_"`
	}
)

func newActRuEventSubscrModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActRuEventSubscrModel {
	return &defaultActRuEventSubscrModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_ru_event_subscr`",
	}
}

func (m *defaultActRuEventSubscrModel) Insert(ctx context.Context, data *ActRuEventSubscr) (sql.Result, error) {
	assetActRuEventSubscrIDKey := fmt.Sprintf("%s%v", cacheAssetActRuEventSubscrIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actRuEventSubscrRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.EVENTTYPE, data.EVENTNAME, data.EXECUTIONID, data.PROCINSTID, data.ACTIVITYID, data.CONFIGURATION, data.CREATED, data.PROCDEFID, data.TENANTID, data.SUBSCOPEID, data.SCOPEID, data.SCOPEDEFINITIONID, data.SCOPETYPE)
	}, assetActRuEventSubscrIDKey)
	return ret, err
}

func (m *defaultActRuEventSubscrModel) FindOne(ctx context.Context, iD string) (*ActRuEventSubscr, error) {
	assetActRuEventSubscrIDKey := fmt.Sprintf("%s%v", cacheAssetActRuEventSubscrIDPrefix, iD)
	var resp ActRuEventSubscr
	err := m.QueryRowCtx(ctx, &resp, assetActRuEventSubscrIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actRuEventSubscrRows, m.table)
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

func (m *defaultActRuEventSubscrModel) Update(ctx context.Context, data *ActRuEventSubscr) error {
	assetActRuEventSubscrIDKey := fmt.Sprintf("%s%v", cacheAssetActRuEventSubscrIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actRuEventSubscrRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.EVENTTYPE, data.EVENTNAME, data.EXECUTIONID, data.PROCINSTID, data.ACTIVITYID, data.CONFIGURATION, data.CREATED, data.PROCDEFID, data.TENANTID, data.SUBSCOPEID, data.SCOPEID, data.SCOPEDEFINITIONID, data.SCOPETYPE, data.ID)
	}, assetActRuEventSubscrIDKey)
	return err
}

func (m *defaultActRuEventSubscrModel) Delete(ctx context.Context, iD string) error {
	assetActRuEventSubscrIDKey := fmt.Sprintf("%s%v", cacheAssetActRuEventSubscrIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActRuEventSubscrIDKey)
	return err
}

func (m *defaultActRuEventSubscrModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActRuEventSubscrIDPrefix, primary)
}

func (m *defaultActRuEventSubscrModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actRuEventSubscrRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActRuEventSubscrModel) tableName() string {
	return m.table
}
