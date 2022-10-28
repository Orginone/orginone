package marketmenu

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAppMenuListLogic {
	return GetAppMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppMenuListLogic) GetAppMenuList(req types.GetAppMenuListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
