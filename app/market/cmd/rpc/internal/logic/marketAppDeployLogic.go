package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/tools/date"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketAppDeployLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarketAppDeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketAppDeployLogic {
	return &MarketAppDeployLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用部署
func (l *MarketAppDeployLogic) MarketAppDeploy(in *market.AppDeployReq) (*market.CommonRpcRes, error) {
	_, err := l.svcCtx.MarketStore.AsMarketApp.UpdateOneID(in.AppId).SetStatus(3).SetApplyTime(date.Now()).Save(l.ctx)
	return &market.CommonRpcRes{Json: "app deploy success."}, err
}
