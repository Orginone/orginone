package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketroledistribution"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppRoleDistribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppRoleDistribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppRoleDistribeLogic {
	return &AppRoleDistribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用权限分配
func (l *AppRoleDistribeLogic) AppRoleDistribe(in *market.AppRoleDistribeRpcReq) (*market.CommonRpcRes, error) {
	tx, err := l.svcCtx.MarketStore.Tx(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	creates := make([]*schema.AsMarketRoleDistributionCreate, 0)
	for _, item := range in.Data {
		if item.RoleId < 1 {
			continue
		}
		_, err = tx.AsMarketRoleDistribution.Update().Where(asmarketroledistribution.IsDeleted(0),
			asmarketroledistribution.RoleID(item.RoleId), asmarketroledistribution.TenantCode(in.TenantCode)).
			SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
		for _, agencyId := range item.AgencyIds {
			creates = append(creates, tx.AsMarketRoleDistribution.Create().SetAgencyID(agencyId).SetTenantCode(in.TenantCode).SetRoleID(item.RoleId))
		}
		for _, jobId := range item.JobIds {
			creates = append(creates, tx.AsMarketRoleDistribution.Create().SetJobID(jobId).SetTenantCode(in.TenantCode).SetRoleID(item.RoleId))
		}
		for _, userId := range item.UserIds {
			creates = append(creates, tx.AsMarketRoleDistribution.Create().SetUserID(userId).SetTenantCode(in.TenantCode).SetRoleID(item.RoleId))
		}
	}
	_, err = tx.AsMarketRoleDistribution.CreateBulk(creates...).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return market.Result("true", err)
}
