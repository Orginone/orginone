package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asappgroupdistributiondata"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/asmarketappgroupdistributionrelation"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketroledistribution"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAppUnsubscribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAppUnsubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAppUnsubscribeLogic {
	return &GroupAppUnsubscribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 集团应用退订
func (l *GroupAppUnsubscribeLogic) GroupAppUnsubscribe(in *market.GroupAppUnsubscribeReq) (*market.CommonRpcRes, error) {
	//开启事务
	tx, err := l.svcCtx.MarketStore.Tx(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppGroupDistributionRelation.Update().Where(asmarketappgroupdistributionrelation.IsDeleted(0),
		asmarketappgroupdistributionrelation.AppIDIn(in.AppIds...), asmarketappgroupdistributionrelation.GroupID(in.GroupId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppGroupDistribution.Update().Where(asmarketappgroupdistribution.IsDeleted(0),
		asmarketappgroupdistribution.AppIDIn(in.AppIds...), asmarketappgroupdistribution.GroupID(in.GroupId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	exist, err := tx.AsMarketAppGroupDistribution.Query().Where(asmarketappgroupdistribution.IsDeleted(0),
		asmarketappgroupdistribution.AppIDIn(in.AppIds...), asmarketappgroupdistribution.TenantID(in.TenantCode)).Exist(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	//注意是不存在
	if !exist {
		_, err = tx.AsMarketRoleDistribution.Update().Where(asmarketroledistribution.IsDeleted(0),
			asmarketroledistribution.TenantCode(in.TenantCode), asmarketroledistribution.HasRolexWith(asmarketapprole.AppIDIn(in.AppIds...),
				asmarketapprole.IsDeleted(0))).SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
	}
	_, err = tx.AsAppGroupDistributionData.Update().Where(asappgroupdistributiondata.IsDeleted(0),
		asappgroupdistributiondata.AppIDIn(in.AppIds...), asappgroupdistributiondata.GroupID(in.GroupId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppPurchase.Update().Where(asmarketapppurchase.IsDeleted(0),
		asmarketapppurchase.AppIDIn(in.AppIds...), asmarketapppurchase.GroupID(in.GroupId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return &market.CommonRpcRes{}, err
}
