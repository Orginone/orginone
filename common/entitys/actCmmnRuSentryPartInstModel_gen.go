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
	actCmmnRuSentryPartInstFieldNames          = builder.RawFieldNames(&ActCmmnRuSentryPartInst{})
	actCmmnRuSentryPartInstRows                = strings.Join(actCmmnRuSentryPartInstFieldNames, ",")
	actCmmnRuSentryPartInstRowsExpectAutoSet   = strings.Join(stringx.Remove(actCmmnRuSentryPartInstFieldNames, "`create_time`", "`update_time`"), ",")
	actCmmnRuSentryPartInstRowsWithPlaceHolder = strings.Join(stringx.Remove(actCmmnRuSentryPartInstFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActCmmnRuSentryPartInstIDPrefix = "cache:asset:actCmmnRuSentryPartInst:iD:"
)

type (
	actCmmnRuSentryPartInstModel interface {
		Insert(ctx context.Context, data *ActCmmnRuSentryPartInst) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActCmmnRuSentryPartInst, error)
		Update(ctx context.Context, data *ActCmmnRuSentryPartInst) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActCmmnRuSentryPartInstModel struct {
		sqlc.CachedConn
		table string
	}

	ActCmmnRuSentryPartInst struct {
		ID             string         `db:"ID_"`
		REV            int64          `db:"REV_"`
		CASEDEFID      sql.NullString `db:"CASE_DEF_ID_"`
		CASEINSTID     sql.NullString `db:"CASE_INST_ID_"`
		PLANITEMINSTID sql.NullString `db:"PLAN_ITEM_INST_ID_"`
		ONPARTID       sql.NullString `db:"ON_PART_ID_"`
		IFPARTID       sql.NullString `db:"IF_PART_ID_"`
		TIMESTAMP      sql.NullTime   `db:"TIME_STAMP_"`
	}
)

func newActCmmnRuSentryPartInstModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActCmmnRuSentryPartInstModel {
	return &defaultActCmmnRuSentryPartInstModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_cmmn_ru_sentry_part_inst`",
	}
}

func (m *defaultActCmmnRuSentryPartInstModel) Insert(ctx context.Context, data *ActCmmnRuSentryPartInst) (sql.Result, error) {
	assetActCmmnRuSentryPartInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuSentryPartInstIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, actCmmnRuSentryPartInstRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.CASEDEFID, data.CASEINSTID, data.PLANITEMINSTID, data.ONPARTID, data.IFPARTID, data.TIMESTAMP)
	}, assetActCmmnRuSentryPartInstIDKey)
	return ret, err
}

func (m *defaultActCmmnRuSentryPartInstModel) FindOne(ctx context.Context, iD string) (*ActCmmnRuSentryPartInst, error) {
	assetActCmmnRuSentryPartInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuSentryPartInstIDPrefix, iD)
	var resp ActCmmnRuSentryPartInst
	err := m.QueryRowCtx(ctx, &resp, assetActCmmnRuSentryPartInstIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnRuSentryPartInstRows, m.table)
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

func (m *defaultActCmmnRuSentryPartInstModel) Update(ctx context.Context, data *ActCmmnRuSentryPartInst) error {
	assetActCmmnRuSentryPartInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuSentryPartInstIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actCmmnRuSentryPartInstRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.CASEDEFID, data.CASEINSTID, data.PLANITEMINSTID, data.ONPARTID, data.IFPARTID, data.TIMESTAMP, data.ID)
	}, assetActCmmnRuSentryPartInstIDKey)
	return err
}

func (m *defaultActCmmnRuSentryPartInstModel) Delete(ctx context.Context, iD string) error {
	assetActCmmnRuSentryPartInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuSentryPartInstIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActCmmnRuSentryPartInstIDKey)
	return err
}

func (m *defaultActCmmnRuSentryPartInstModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActCmmnRuSentryPartInstIDPrefix, primary)
}

func (m *defaultActCmmnRuSentryPartInstModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnRuSentryPartInstRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActCmmnRuSentryPartInstModel) tableName() string {
	return m.table
}