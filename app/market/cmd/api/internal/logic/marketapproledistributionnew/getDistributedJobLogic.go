package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributedJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributedJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributedJobLogic {
	return GetDistributedJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributedJobLogic) GetDistributedJob(req types.GetDistributedJobReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
