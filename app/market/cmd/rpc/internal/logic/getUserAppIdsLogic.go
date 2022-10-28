package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAppIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAppIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAppIdsLogic {
	return &GetUserAppIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户使用的应用Ids
func (l *GetUserAppIdsLogic) GetUserAppIds(in *market.GetUserAppIdsReq) (*market.CommonRpcRes, error) {
	userEntity, err := l.svcCtx.MarketStore.AsUser.Query().Where(asuser.IsDeleted(0), asuser.ID(in.UserId)).WithAgencys().WithJobs().First(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	jobIds := make([]int64, 1)
	jobIds[0] = -10000
	agencyIds := make([]int64, 1)
	agencyIds[0] = -10000
	for _, v := range userEntity.Edges.Jobs {
		if v.IsDeleted == 0 {
			jobIds = append(jobIds, v.ID)
		}
	}
	for _, v := range userEntity.Edges.Agencys {
		if v.IsDeleted == 0 {
			agencyIds = append(agencyIds, v.ID)
		}
	}
	return market.JsonResult(l.svcCtx.MarketStore.AsMarketRoleDistribution.Query().
		Where(asmarketroledistribution.IsDeleted(0), asmarketroledistribution.Or(asmarketroledistribution.UserID(in.UserId),
			asmarketroledistribution.AgencyIDIn(agencyIds...), asmarketroledistribution.JobIDIn(jobIds...))).
		QueryRolex().Where(asmarketapprole.IsDeleted(0)).Select(asmarketapprole.FieldAppID).Int64s(l.ctx))
}
