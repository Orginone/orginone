package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributionAgencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributionAgencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributionAgencyLogic {
	return GetDistributionAgencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributionAgencyLogic) GetDistributionAgency(req types.GetDistributionAgencyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
