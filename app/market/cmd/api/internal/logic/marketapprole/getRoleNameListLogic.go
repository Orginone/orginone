package marketapprole

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleNameListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleNameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRoleNameListLogic {
	return GetRoleNameListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleNameListLogic) GetRoleNameList(req types.GetRoleNameListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
