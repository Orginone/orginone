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
	actHiIdentitylinkFieldNames          = builder.RawFieldNames(&ActHiIdentitylink{})
	actHiIdentitylinkRows                = strings.Join(actHiIdentitylinkFieldNames, ",")
	actHiIdentitylinkRowsExpectAutoSet   = strings.Join(stringx.Remove(actHiIdentitylinkFieldNames, "`create_time`", "`update_time`"), ",")
	actHiIdentitylinkRowsWithPlaceHolder = strings.Join(stringx.Remove(actHiIdentitylinkFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActHiIdentitylinkIDPrefix = "cache:asset:actHiIdentitylink:iD:"
)

type (
	actHiIdentitylinkModel interface {
		Insert(ctx context.Context, data *ActHiIdentitylink) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActHiIdentitylink, error)
		Update(ctx context.Context, data *ActHiIdentitylink) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActHiIdentitylinkModel struct {
		sqlc.CachedConn
		table string
	}

	ActHiIdentitylink struct {
		ID                string         `db:"ID_"`
		GROUPID           sql.NullString `db:"GROUP_ID_"`
		TYPE              sql.NullString `db:"TYPE_"`
		USERID            sql.NullString `db:"USER_ID_"`
		TASKID            sql.NullString `db:"TASK_ID_"`
		CREATETIME        sql.NullTime   `db:"CREATE_TIME_"`
		PROCINSTID        sql.NullString `db:"PROC_INST_ID_"`
		SCOPEID           sql.NullString `db:"SCOPE_ID_"`
		SCOPETYPE         sql.NullString `db:"SCOPE_TYPE_"`
		SCOPEDEFINITIONID sql.NullString `db:"SCOPE_DEFINITION_ID_"`
		SUBSCOPEID        sql.NullString `db:"SUB_SCOPE_ID_"`
	}
)

func newActHiIdentitylinkModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActHiIdentitylinkModel {
	return &defaultActHiIdentitylinkModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_hi_identitylink`",
	}
}

func (m *defaultActHiIdentitylinkModel) Insert(ctx context.Context, data *ActHiIdentitylink) (sql.Result, error) {
	assetActHiIdentitylinkIDKey := fmt.Sprintf("%s%v", cacheAssetActHiIdentitylinkIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actHiIdentitylinkRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.GROUPID, data.TYPE, data.USERID, data.TASKID, data.CREATETIME, data.PROCINSTID, data.SCOPEID, data.SCOPETYPE, data.SCOPEDEFINITIONID, data.SUBSCOPEID)
	}, assetActHiIdentitylinkIDKey)
	return ret, err
}

func (m *defaultActHiIdentitylinkModel) FindOne(ctx context.Context, iD string) (*ActHiIdentitylink, error) {
	assetActHiIdentitylinkIDKey := fmt.Sprintf("%s%v", cacheAssetActHiIdentitylinkIDPrefix, iD)
	var resp ActHiIdentitylink
	err := m.QueryRowCtx(ctx, &resp, assetActHiIdentitylinkIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actHiIdentitylinkRows, m.table)
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

func (m *defaultActHiIdentitylinkModel) Update(ctx context.Context, data *ActHiIdentitylink) error {
	assetActHiIdentitylinkIDKey := fmt.Sprintf("%s%v", cacheAssetActHiIdentitylinkIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actHiIdentitylinkRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.GROUPID, data.TYPE, data.USERID, data.TASKID, data.CREATETIME, data.PROCINSTID, data.SCOPEID, data.SCOPETYPE, data.SCOPEDEFINITIONID, data.SUBSCOPEID, data.ID)
	}, assetActHiIdentitylinkIDKey)
	return err
}

func (m *defaultActHiIdentitylinkModel) Delete(ctx context.Context, iD string) error {
	assetActHiIdentitylinkIDKey := fmt.Sprintf("%s%v", cacheAssetActHiIdentitylinkIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActHiIdentitylinkIDKey)
	return err
}

func (m *defaultActHiIdentitylinkModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActHiIdentitylinkIDPrefix, primary)
}

func (m *defaultActHiIdentitylinkModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actHiIdentitylinkRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActHiIdentitylinkModel) tableName() string {
	return m.table
}
