package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappnotice"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMarkAppNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveMarkAppNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMarkAppNoticeListLogic {
	return &RemoveMarkAppNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据ID列表移除应用通知
func (l *RemoveMarkAppNoticeListLogic) RemoveMarkAppNoticeList(in *market.MarkAppNoticeIdsReq) (*market.CommonRpcRes, error) {
	return market.JsonResult(l.svcCtx.MarketStore.AsMarketAppNotice.Update().Where(asmarketappnotice.IDIn(in.NotoceIds...)).SetIsDeleted(1).Save(l.ctx))
}
