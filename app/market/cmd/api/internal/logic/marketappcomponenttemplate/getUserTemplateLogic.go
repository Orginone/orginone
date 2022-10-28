package marketappcomponenttemplate

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserTemplateLogic {
	return GetUserTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserTemplateLogic) GetUserTemplate(req types.GetUserTemplateReq) (resp *types.CommonResponse, err error) {
	userId, _, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.PageResult(l.svcCtx.MarketRpc.GetUserTemplates(l.ctx, &market.GetUserTemplateReq{
		UserId: userId,
		Page: &market.PageRequest{
			Limit:  req.Size,
			Filter: req.Name,
			Offset: (req.Current - 1) * req.Size,
		},
	}))
}
