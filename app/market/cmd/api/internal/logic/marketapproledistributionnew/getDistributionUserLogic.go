package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributionUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributionUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributionUserLogic {
	return GetDistributionUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributionUserLogic) GetDistributionUser(req types.GetDistributionUserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
