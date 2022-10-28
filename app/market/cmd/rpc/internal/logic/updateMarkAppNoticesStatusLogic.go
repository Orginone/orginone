package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/tools/date"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMarkAppNoticesStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMarkAppNoticesStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMarkAppNoticesStatusLogic {
	return &UpdateMarkAppNoticesStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新应用通知状态
func (l *UpdateMarkAppNoticesStatusLogic) UpdateMarkAppNoticesStatus(in *market.UpdateMarkAppNoticesStatusReq) (*market.CommonRpcRes, error) {
	update := l.svcCtx.MarketStore.AsMarketAppNotice.UpdateOneID(in.NotoceId)
	if in.NoticeStatus == 1 {
		update.SetNoticeReleaseTime(date.Now())
	}
	return market.JsonResult(update.SetNoticeReleaseStatus(in.NoticeStatus).SetIsDeleted(in.IsDeleted).Save(l.ctx))

}
