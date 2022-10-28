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
	actProcdefInfoFieldNames          = builder.RawFieldNames(&ActProcdefInfo{})
	actProcdefInfoRows                = strings.Join(actProcdefInfoFieldNames, ",")
	actProcdefInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(actProcdefInfoFieldNames, "`create_time`", "`update_time`"), ",")
	actProcdefInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(actProcdefInfoFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActProcdefInfoIDPrefix        = "cache:asset:actProcdefInfo:iD:"
	cacheAssetActProcdefInfoPROCDEFIDPrefix = "cache:asset:actProcdefInfo:pROCDEFID:"
)

type (
	actProcdefInfoModel interface {
		Insert(ctx context.Context, data *ActProcdefInfo) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActProcdefInfo, error)
		FindOneByPROCDEFID(ctx context.Context, pROCDEFID string) (*ActProcdefInfo, error)
		Update(ctx context.Context, data *ActProcdefInfo) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActProcdefInfoModel struct {
		sqlc.CachedConn
		table string
	}

	ActProcdefInfo struct {
		ID         string         `db:"ID_"`
		PROCDEFID  string         `db:"PROC_DEF_ID_"`
		REV        sql.NullInt64  `db:"REV_"`
		INFOJSONID sql.NullString `db:"INFO_JSON_ID_"`
	}
)

func newActProcdefInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActProcdefInfoModel {
	return &defaultActProcdefInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_procdef_info`",
	}
}

func (m *defaultActProcdefInfoModel) Insert(ctx context.Context, data *ActProcdefInfo) (sql.Result, error) {
	assetActProcdefInfoIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoIDPrefix, data.ID)
	assetActProcdefInfoPROCDEFIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoPROCDEFIDPrefix, data.PROCDEFID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, actProcdefInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.PROCDEFID, data.REV, data.INFOJSONID)
	}, assetActProcdefInfoIDKey, assetActProcdefInfoPROCDEFIDKey)
	return ret, err
}

func (m *defaultActProcdefInfoModel) FindOne(ctx context.Context, iD string) (*ActProcdefInfo, error) {
	assetActProcdefInfoIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoIDPrefix, iD)
	var resp ActProcdefInfo
	err := m.QueryRowCtx(ctx, &resp, assetActProcdefInfoIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actProcdefInfoRows, m.table)
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

func (m *defaultActProcdefInfoModel) FindOneByPROCDEFID(ctx context.Context, pROCDEFID string) (*ActProcdefInfo, error) {
	assetActProcdefInfoPROCDEFIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoPROCDEFIDPrefix, pROCDEFID)
	var resp ActProcdefInfo
	err := m.QueryRowIndexCtx(ctx, &resp, assetActProcdefInfoPROCDEFIDKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `PROC_DEF_ID_` = ? limit 1", actProcdefInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, pROCDEFID); err != nil {
			return nil, err
		}
		return resp.ID, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultActProcdefInfoModel) Update(ctx context.Context, data *ActProcdefInfo) error {
	assetActProcdefInfoIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoIDPrefix, data.ID)
	assetActProcdefInfoPROCDEFIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoPROCDEFIDPrefix, data.PROCDEFID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actProcdefInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.PROCDEFID, data.REV, data.INFOJSONID, data.ID)
	}, assetActProcdefInfoIDKey, assetActProcdefInfoPROCDEFIDKey)
	return err
}

func (m *defaultActProcdefInfoModel) Delete(ctx context.Context, iD string) error {
	data, err := m.FindOne(ctx, iD)
	if err != nil {
		return err
	}

	assetActProcdefInfoIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoIDPrefix, iD)
	assetActProcdefInfoPROCDEFIDKey := fmt.Sprintf("%s%v", cacheAssetActProcdefInfoPROCDEFIDPrefix, data.PROCDEFID)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActProcdefInfoIDKey, assetActProcdefInfoPROCDEFIDKey)
	return err
}

func (m *defaultActProcdefInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActProcdefInfoIDPrefix, primary)
}

func (m *defaultActProcdefInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actProcdefInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActProcdefInfoModel) tableName() string {
	return m.table
}
