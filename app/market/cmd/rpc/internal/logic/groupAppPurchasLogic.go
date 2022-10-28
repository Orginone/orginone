package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAppPurchasLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAppPurchasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAppPurchasLogic {
	return &GroupAppPurchasLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 集团获取应用
func (l *GroupAppPurchasLogic) GroupAppPurchas(in *market.GroupAppPurchasReq) (*market.CommonRpcRes, error) {
	for _, groupId := range in.GroupIds {
		_, err := l.svcCtx.MarketStore.AsMarketAppPurchase.Create().SetAppID(in.AppId).SetGroupID(groupId).SetUseStatus(1).Save(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
	}
	return &market.CommonRpcRes{Json: "true"}, nil
}
