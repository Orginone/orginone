package menus

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetmenusbyroleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetmenusbyroleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetmenusbyroleIdLogic {
	return GetmenusbyroleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetmenusbyroleIdLogic) GetmenusbyroleId(req types.GetMenusByRoleIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
