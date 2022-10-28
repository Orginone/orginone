package marketappnotice

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListReq7) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.FindMarkAppNoticeList(l.ctx, &market.FindMarkAppNoticeListReq{
		Page: &market.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		GroupOrUnit: req.GroupOrUnit,
		NoticeTitle: req.NoticeTitle,
	}))
}
