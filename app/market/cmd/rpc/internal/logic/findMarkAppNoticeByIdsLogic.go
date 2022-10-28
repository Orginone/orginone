package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappnotice"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMarkAppNoticeByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMarkAppNoticeByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMarkAppNoticeByIdsLogic {
	return &FindMarkAppNoticeByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据ID列表获取应用通知
func (l *FindMarkAppNoticeByIdsLogic) FindMarkAppNoticeByIds(in *market.MarkAppNoticeIdsReq) (*market.CommonRpcRes, error) {
	return market.JsonResult(l.svcCtx.MarketStore.AsMarketAppNotice.Query().Where(asmarketappnotice.IDIn(in.NotoceIds...)).All(l.ctx))
}
