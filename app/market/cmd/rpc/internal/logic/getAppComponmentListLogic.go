package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappcomponent"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppComponmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppComponmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppComponmentListLogic {
	return &GetAppComponmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用组件
func (l *GetAppComponmentListLogic) GetAppComponmentList(in *market.AppLinkReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppComponent.Query().Where(asmarketappcomponent.IsDeleted(0),
		asmarketappcomponent.AppID(in.AppId))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	componmentEntitys, err := query.Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return market.PageResult(in.Page, total, componmentEntitys, err)
}
