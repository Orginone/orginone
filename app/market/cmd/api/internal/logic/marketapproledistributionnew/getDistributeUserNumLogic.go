package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributeUserNumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributeUserNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributeUserNumLogic {
	return GetDistributeUserNumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributeUserNumLogic) GetDistributeUserNum(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}