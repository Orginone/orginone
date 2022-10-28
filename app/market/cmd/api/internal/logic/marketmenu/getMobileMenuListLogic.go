package marketmenu

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMobileMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMobileMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMobileMenuListLogic {
	return GetMobileMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMobileMenuListLogic) GetMobileMenuList(req types.GetMobileMenuListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
