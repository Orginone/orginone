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
	tSysPostFieldNames          = builder.RawFieldNames(&TSysPost{})
	tSysPostRows                = strings.Join(tSysPostFieldNames, ",")
	tSysPostRowsExpectAutoSet   = strings.Join(stringx.Remove(tSysPostFieldNames, "`create_time`", "`update_time`"), ",")
	tSysPostRowsWithPlaceHolder = strings.Join(stringx.Remove(tSysPostFieldNames, "`POST_ID`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetTSysPostPOSTIDPrefix = "cache:asset:tSysPost:pOSTID:"
)

type (
	tSysPostModel interface {
		Insert(ctx context.Context, data *TSysPost) (sql.Result, error)
		FindOne(ctx context.Context, pOSTID string) (*TSysPost, error)
		Update(ctx context.Context, data *TSysPost) error
		Delete(ctx context.Context, pOSTID string) error
	}

	defaultTSysPostModel struct {
		sqlc.CachedConn
		table string
	}

	TSysPost struct {
		POSTID     string         `db:"POST_ID"`     // 岗位ID
		POSTNAME   string         `db:"POST_NAME"`   // 岗位名称
		POSTSTATUS sql.NullString `db:"POST_STATUS"` // 状态
		SORTNO     sql.NullInt64  `db:"SORT_NO"`     // 排序号
		REMARK     sql.NullString `db:"REMARK"`      // 备注
		CREATEBY   sql.NullString `db:"CREATE_BY"`   // 创建人
		CREATEDATE sql.NullTime   `db:"CREATE_DATE"` // 创建日期
		CREATETIME sql.NullTime   `db:"CREATE_TIME"` // 创建时间
		UPDATEBY   sql.NullString `db:"UPDATE_BY"`   // 修改人
		UPDATEDATE sql.NullTime   `db:"UPDATE_DATE"` // 修改日期
		UPDATETIME sql.NullTime   `db:"UPDATE_TIME"` // 修改时间
	}
)

func newTSysPostModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTSysPostModel {
	return &defaultTSysPostModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`t_sys_post`",
	}
}

func (m *defaultTSysPostModel) Insert(ctx context.Context, data *TSysPost) (sql.Result, error) {
	assetTSysPostPOSTIDKey := fmt.Sprintf("%s%v", cacheAssetTSysPostPOSTIDPrefix, data.POSTID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tSysPostRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.POSTID, data.POSTNAME, data.POSTSTATUS, data.SORTNO, data.REMARK, data.CREATEBY, data.CREATEDATE, data.CREATETIME, data.UPDATEBY, data.UPDATEDATE, data.UPDATETIME)
	}, assetTSysPostPOSTIDKey)
	return ret, err
}

func (m *defaultTSysPostModel) FindOne(ctx context.Context, pOSTID string) (*TSysPost, error) {
	assetTSysPostPOSTIDKey := fmt.Sprintf("%s%v", cacheAssetTSysPostPOSTIDPrefix, pOSTID)
	var resp TSysPost
	err := m.QueryRowCtx(ctx, &resp, assetTSysPostPOSTIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `POST_ID` = ? limit 1", tSysPostRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, pOSTID)
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

func (m *defaultTSysPostModel) Update(ctx context.Context, data *TSysPost) error {
	assetTSysPostPOSTIDKey := fmt.Sprintf("%s%v", cacheAssetTSysPostPOSTIDPrefix, data.POSTID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `POST_ID` = ?", m.table, tSysPostRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.POSTNAME, data.POSTSTATUS, data.SORTNO, data.REMARK, data.CREATEBY, data.CREATEDATE, data.CREATETIME, data.UPDATEBY, data.UPDATEDATE, data.UPDATETIME, data.POSTID)
	}, assetTSysPostPOSTIDKey)
	return err
}

func (m *defaultTSysPostModel) Delete(ctx context.Context, pOSTID string) error {
	assetTSysPostPOSTIDKey := fmt.Sprintf("%s%v", cacheAssetTSysPostPOSTIDPrefix, pOSTID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `POST_ID` = ?", m.table)
		return conn.ExecCtx(ctx, query, pOSTID)
	}, assetTSysPostPOSTIDKey)
	return err
}

func (m *defaultTSysPostModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetTSysPostPOSTIDPrefix, primary)
}

func (m *defaultTSysPostModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `POST_ID` = ? limit 1", tSysPostRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTSysPostModel) tableName() string {
	return m.table
}
