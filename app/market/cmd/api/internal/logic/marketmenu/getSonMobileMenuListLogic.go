package marketmenu

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSonMobileMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSonMobileMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSonMobileMenuListLogic {
	return GetSonMobileMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSonMobileMenuListLogic) GetSonMobileMenuList(req types.GetSonMobileMenuListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
