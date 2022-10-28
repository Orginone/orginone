package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributionPersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributionPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributionPersonLogic {
	return GetDistributionPersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributionPersonLogic) GetDistributionPerson(req types.GetDistributionPersonReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
