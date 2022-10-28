package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRoleByUserIdLogic {
	return GetRoleByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleByUserIdLogic) GetRoleByUserId(req types.GetRoleByUserIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
