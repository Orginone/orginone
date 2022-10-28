package marketapp

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitFirstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitFirstLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitFirstLogic {
	return SubmitFirstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitFirstLogic) SubmitFirst(req []byte) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.MarketRpc.SubmitMarketApp(l.ctx, &market.CommonRpcReq{Json: string(req)})
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	return types.Successful(true)
}
