package marketappnotice

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(Ids []string) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.RemoveMarkAppNoticeList(l.ctx, &market.MarkAppNoticeIdsReq{
		NotoceIds: tools.StrArrayToNumArray(Ids),
	}))
}
