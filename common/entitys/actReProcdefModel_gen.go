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
	actReProcdefFieldNames          = builder.RawFieldNames(&ActReProcdef{})
	actReProcdefRows                = strings.Join(actReProcdefFieldNames, ",")
	actReProcdefRowsExpectAutoSet   = strings.Join(stringx.Remove(actReProcdefFieldNames, "`create_time`", "`update_time`"), ",")
	actReProcdefRowsWithPlaceHolder = strings.Join(stringx.Remove(actReProcdefFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActReProcdefIDPrefix                               = "cache:asset:actReProcdef:iD:"
	cacheAssetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDPrefix = "cache:asset:actReProcdef:kEY:vERSION:dERIVEDVERSION:tENANTID:"
)

type (
	actReProcdefModel interface {
		Insert(ctx context.Context, data *ActReProcdef) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActReProcdef, error)
		FindOneByKEYVERSIONDERIVEDVERSIONTENANTID(ctx context.Context, kEY string, vERSION int64, dERIVEDVERSION int64, tENANTID string) (*ActReProcdef, error)
		Update(ctx context.Context, data *ActReProcdef) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActReProcdefModel struct {
		sqlc.CachedConn
		table string
	}

	ActReProcdef struct {
		ID                   string         `db:"ID_"`
		REV                  sql.NullInt64  `db:"REV_"`
		CATEGORY             sql.NullString `db:"CATEGORY_"`
		NAME                 sql.NullString `db:"NAME_"`
		KEY                  string         `db:"KEY_"`
		VERSION              int64          `db:"VERSION_"`
		DEPLOYMENTID         sql.NullString `db:"DEPLOYMENT_ID_"`
		RESOURCENAME         sql.NullString `db:"RESOURCE_NAME_"`
		DGRMRESOURCENAME     sql.NullString `db:"DGRM_RESOURCE_NAME_"`
		DESCRIPTION          sql.NullString `db:"DESCRIPTION_"`
		HASSTARTFORMKEY      sql.NullInt64  `db:"HAS_START_FORM_KEY_"`
		HASGRAPHICALNOTATION sql.NullInt64  `db:"HAS_GRAPHICAL_NOTATION_"`
		SUSPENSIONSTATE      sql.NullInt64  `db:"SUSPENSION_STATE_"`
		TENANTID             string         `db:"TENANT_ID_"`
		ENGINEVERSION        sql.NullString `db:"ENGINE_VERSION_"`
		DERIVEDFROM          sql.NullString `db:"DERIVED_FROM_"`
		DERIVEDFROMROOT      sql.NullString `db:"DERIVED_FROM_ROOT_"`
		DERIVEDVERSION       int64          `db:"DERIVED_VERSION_"`
	}
)

func newActReProcdefModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActReProcdefModel {
	return &defaultActReProcdefModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_re_procdef`",
	}
}

func (m *defaultActReProcdefModel) Insert(ctx context.Context, data *ActReProcdef) (sql.Result, error) {
	assetActReProcdefIDKey := fmt.Sprintf("%s%v", cacheAssetActReProcdefIDPrefix, data.ID)
	assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheAssetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDPrefix, data.KEY, data.VERSION, data.DERIVEDVERSION, data.TENANTID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actReProcdefRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.CATEGORY, data.NAME, data.KEY, data.VERSION, data.DEPLOYMENTID, data.RESOURCENAME, data.DGRMRESOURCENAME, data.DESCRIPTION, data.HASSTARTFORMKEY, data.HASGRAPHICALNOTATION, data.SUSPENSIONSTATE, data.TENANTID, data.ENGINEVERSION, data.DERIVEDFROM, data.DERIVEDFROMROOT, data.DERIVEDVERSION)
	}, assetActReProcdefIDKey, assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey)
	return ret, err
}

func (m *defaultActReProcdefModel) FindOne(ctx context.Context, iD string) (*ActReProcdef, error) {
	assetActReProcdefIDKey := fmt.Sprintf("%s%v", cacheAssetActReProcdefIDPrefix, iD)
	var resp ActReProcdef
	err := m.QueryRowCtx(ctx, &resp, assetActReProcdefIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actReProcdefRows, m.table)
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

func (m *defaultActReProcdefModel) FindOneByKEYVERSIONDERIVEDVERSIONTENANTID(ctx context.Context, kEY string, vERSION int64, dERIVEDVERSION int64, tENANTID string) (*ActReProcdef, error) {
	assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheAssetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDPrefix, kEY, vERSION, dERIVEDVERSION, tENANTID)
	var resp ActReProcdef
	err := m.QueryRowIndexCtx(ctx, &resp, assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `KEY_` = ? and `VERSION_` = ? and `DERIVED_VERSION_` = ? and `TENANT_ID_` = ? limit 1", actReProcdefRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, kEY, vERSION, dERIVEDVERSION, tENANTID); err != nil {
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

func (m *defaultActReProcdefModel) Update(ctx context.Context, data *ActReProcdef) error {
	assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheAssetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDPrefix, data.KEY, data.VERSION, data.DERIVEDVERSION, data.TENANTID)
	assetActReProcdefIDKey := fmt.Sprintf("%s%v", cacheAssetActReProcdefIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actReProcdefRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.CATEGORY, data.NAME, data.KEY, data.VERSION, data.DEPLOYMENTID, data.RESOURCENAME, data.DGRMRESOURCENAME, data.DESCRIPTION, data.HASSTARTFORMKEY, data.HASGRAPHICALNOTATION, data.SUSPENSIONSTATE, data.TENANTID, data.ENGINEVERSION, data.DERIVEDFROM, data.DERIVEDFROMROOT, data.DERIVEDVERSION, data.ID)
	}, assetActReProcdefIDKey, assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey)
	return err
}

func (m *defaultActReProcdefModel) Delete(ctx context.Context, iD string) error {
	data, err := m.FindOne(ctx, iD)
	if err != nil {
		return err
	}

	assetActReProcdefIDKey := fmt.Sprintf("%s%v", cacheAssetActReProcdefIDPrefix, iD)
	assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheAssetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDPrefix, data.KEY, data.VERSION, data.DERIVEDVERSION, data.TENANTID)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActReProcdefIDKey, assetActReProcdefKEYVERSIONDERIVEDVERSIONTENANTIDKey)
	return err
}

func (m *defaultActReProcdefModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActReProcdefIDPrefix, primary)
}

func (m *defaultActReProcdefModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actReProcdefRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActReProcdefModel) tableName() string {
	return m.table
}
