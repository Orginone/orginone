package entitys

import (
	"context"
	"fmt"
	"orginone/common/tools"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AsInnerAgencyModel = (*customAsInnerAgencyModel)(nil)

type (
	// AsInnerAgencyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAsInnerAgencyModel.
	AsInnerAgencyModel interface {
		asInnerAgencyModel
		FindByAgencyIds(ctx context.Context, agencyIds []int64) ([]*AsInnerAgency, error)
	}

	customAsInnerAgencyModel struct {
		*defaultAsInnerAgencyModel
	}
)

// NewAsInnerAgencyModel returns a model for the database table.
func NewAsInnerAgencyModel(conn sqlx.SqlConn, c cache.CacheConf) AsInnerAgencyModel {
	return &customAsInnerAgencyModel{
		defaultAsInnerAgencyModel: newAsInnerAgencyModel(conn, c),
	}
}

//根据角色Ids查找
func (m *customAsInnerAgencyModel) FindByAgencyIds(ctx context.Context, agencyIds []int64) ([]*AsInnerAgency, error) {
	query, values, err := squirrel.Select(asInnerAgencyRows).From(m.table).Where(fmt.Sprintf("id in %s", tools.SqlInArgsMap(agencyIds))).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*AsInnerAgency
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
