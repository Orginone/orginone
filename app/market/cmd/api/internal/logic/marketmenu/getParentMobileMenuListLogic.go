package marketmenu

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParentMobileMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetParentMobileMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetParentMobileMenuListLogic {
	return GetParentMobileMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetParentMobileMenuListLogic) GetParentMobileMenuList(req types.GetParentMobileMenuListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
