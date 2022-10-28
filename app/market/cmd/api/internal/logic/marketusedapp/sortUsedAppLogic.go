package marketusedapp

import (
	"context"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortUsedAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSortUsedAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) SortUsedAppLogic {
	return SortUsedAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SortUsedAppLogic) SortUsedApp(appIdList []int64) (resp *types.CommonResponse, err error) {
	userId, _, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.CommonResult(l.svcCtx.MarketRpc.SortUsedApp(l.ctx, &market.SortUsedAppReq{
		AppIdList: appIdList,
		UserId:    userId,
	}))
}
