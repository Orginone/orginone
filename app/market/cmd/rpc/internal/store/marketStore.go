package store

import (
	"context"
	"orginone/common"
	"orginone/common/models"
	"orginone/common/schema/astenant"
)

type MarketStore struct {
	*common.DbStore
}

func NewMarketStore(c models.DbStoreConfig) *MarketStore {
	return &MarketStore{
		DbStore: common.NewDbStore(c),
	}
}

//根据租户编号查询租户名称
func (s *MarketStore) GetTenantNameByTenantCode(ctx context.Context, tenantCode string, defaultName string) string {
	tenantNames, err := s.AsTenant.Query().Where(astenant.TenantCode(tenantCode), astenant.IsDeleted(0)).
		Limit(1).Select(astenant.FieldTenantName).Strings(ctx)
	if err == nil && len(tenantNames) > 0 {
		return tenantNames[0]
	}
	return defaultName
}
