package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributedAgencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributedAgencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributedAgencyLogic {
	return GetDistributedAgencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributedAgencyLogic) GetDistributedAgency(req types.GetDistributedAgencyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
