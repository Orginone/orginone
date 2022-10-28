package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppPurchaseGroupMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppPurchaseGroupMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAppPurchaseGroupMapLogic {
	return GetAppPurchaseGroupMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppPurchaseGroupMapLogic) GetAppPurchaseGroupMap(req types.GetAppPurchaseGroupMapReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
