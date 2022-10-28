package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateAllPortalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateAllPortalLogic(ctx context.Context, svcCtx *svc.ServiceContext) SaveOrUpdateAllPortalLogic {
	return SaveOrUpdateAllPortalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateAllPortalLogic) SaveOrUpdateAllPortal(req types.SaveOrUpdateAllPortalReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
