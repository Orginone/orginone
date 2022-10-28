package marketappcomponent

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComponentListByAppIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetComponentListByAppIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetComponentListByAppIdLogic {
	return GetComponentListByAppIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComponentListByAppIdLogic) GetComponentListByAppId(req types.GetComponentListByAppIdReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.MarketRpc.GetAppComponmentList(l.ctx, &market.AppLinkReq{
		AppId: req.AppId,
		Page: &market.PageRequest{
			Limit:  req.Size,
			Offset: (req.Current - 1) * req.Size,
		},
	}))
}
