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
	basedataAssetclassFieldNames          = builder.RawFieldNames(&BasedataAssetclass{})
	basedataAssetclassRows                = strings.Join(basedataAssetclassFieldNames, ",")
	basedataAssetclassRowsExpectAutoSet   = strings.Join(stringx.Remove(basedataAssetclassFieldNames, "`create_time`", "`update_time`"), ",")
	basedataAssetclassRowsWithPlaceHolder = strings.Join(stringx.Remove(basedataAssetclassFieldNames, "`RECID`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBasedataAssetclassRECIDPrefix = "cache:asset:basedataAssetclass:rECID:"
)

type (
	basedataAssetclassModel interface {
		Insert(ctx context.Context, data *BasedataAssetclass) (sql.Result, error)
		FindOne(ctx context.Context, rECID int64) (*BasedataAssetclass, error)
		Update(ctx context.Context, data *BasedataAssetclass) error
		Delete(ctx context.Context, rECID int64) error
	}

	defaultBasedataAssetclassModel struct {
		sqlc.CachedConn
		table string
	}

	BasedataAssetclass struct {
		RECID        int64           `db:"RECID"`
		RECVER       sql.NullInt64   `db:"RECVER"`
		OLDID        sql.NullString  `db:"OLDID"`
		CODE         sql.NullString  `db:"CODE"`
		NAME         sql.NullString  `db:"NAME"`
		SHIFOUMJJD   sql.NullString  `db:"SHIFOUMJJD"`
		ORDER        sql.NullFloat64 `db:"ORDER_"`
		JILIANGDW    sql.NullString  `db:"JILIANGDW"`
		ZICHANDLID   sql.NullString  `db:"ZICHANDLID"`
		SHORTNAME    sql.NullString  `db:"SHORTNAME"`
		LEVEL        sql.NullFloat64 `db:"LEVEL"`
		PARENTID     sql.NullString  `db:"PARENTID"`
		STATE        sql.NullInt64   `db:"STATE"`
		FenLeiGBidId sql.NullString  `db:"fenLeiGBidId"`
		SHIYNX       sql.NullFloat64 `db:"SHIYNX"`
		ZUIGAOSYNX   sql.NullFloat64 `db:"ZUIGAOSYNX"`
		STYLEID      sql.NullInt64   `db:"STYLEID"`
		ORGID        sql.NullString  `db:"ORGID"`
	}
)

func newBasedataAssetclassModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBasedataAssetclassModel {
	return &defaultBasedataAssetclassModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`basedata_assetclass`",
	}
}

func (m *defaultBasedataAssetclassModel) Insert(ctx context.Context, data *BasedataAssetclass) (sql.Result, error) {
	assetBasedataAssetclassRECIDKey := fmt.Sprintf("%s%v", cacheAssetBasedataAssetclassRECIDPrefix, data.RECID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, basedataAssetclassRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RECID, data.RECVER, data.OLDID, data.CODE, data.NAME, data.SHIFOUMJJD, data.ORDER, data.JILIANGDW, data.ZICHANDLID, data.SHORTNAME, data.LEVEL, data.PARENTID, data.STATE, data.FenLeiGBidId, data.SHIYNX, data.ZUIGAOSYNX, data.STYLEID, data.ORGID)
	}, assetBasedataAssetclassRECIDKey)
	return ret, err
}

func (m *defaultBasedataAssetclassModel) FindOne(ctx context.Context, rECID int64) (*BasedataAssetclass, error) {
	assetBasedataAssetclassRECIDKey := fmt.Sprintf("%s%v", cacheAssetBasedataAssetclassRECIDPrefix, rECID)
	var resp BasedataAssetclass
	err := m.QueryRowCtx(ctx, &resp, assetBasedataAssetclassRECIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `RECID` = ? limit 1", basedataAssetclassRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, rECID)
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

func (m *defaultBasedataAssetclassModel) Update(ctx context.Context, data *BasedataAssetclass) error {
	assetBasedataAssetclassRECIDKey := fmt.Sprintf("%s%v", cacheAssetBasedataAssetclassRECIDPrefix, data.RECID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `RECID` = ?", m.table, basedataAssetclassRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RECVER, data.OLDID, data.CODE, data.NAME, data.SHIFOUMJJD, data.ORDER, data.JILIANGDW, data.ZICHANDLID, data.SHORTNAME, data.LEVEL, data.PARENTID, data.STATE, data.FenLeiGBidId, data.SHIYNX, data.ZUIGAOSYNX, data.STYLEID, data.ORGID, data.RECID)
	}, assetBasedataAssetclassRECIDKey)
	return err
}

func (m *defaultBasedataAssetclassModel) Delete(ctx context.Context, rECID int64) error {
	assetBasedataAssetclassRECIDKey := fmt.Sprintf("%s%v", cacheAssetBasedataAssetclassRECIDPrefix, rECID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `RECID` = ?", m.table)
		return conn.ExecCtx(ctx, query, rECID)
	}, assetBasedataAssetclassRECIDKey)
	return err
}

func (m *defaultBasedataAssetclassModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBasedataAssetclassRECIDPrefix, primary)
}

func (m *defaultBasedataAssetclassModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `RECID` = ? limit 1", basedataAssetclassRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBasedataAssetclassModel) tableName() string {
	return m.table
}