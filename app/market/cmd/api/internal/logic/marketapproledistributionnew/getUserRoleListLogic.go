package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserRoleListLogic {
	return GetUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoleListLogic) GetUserRoleList(req types.GetUserRoleListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
