package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDefaultTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDefaultTemplateLogic {
	return GetDefaultTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDefaultTemplateLogic) GetDefaultTemplate(req types.GetDefaultTemplateReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.MarketRpc.GetUserTemplateId(l.ctx, &market.GetUserTemplateIdReq{
		UserId: req.UserId,
	}))
}
