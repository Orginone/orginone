package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetlayergroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetlayergroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetlayergroupLogic {
	return GetlayergroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetlayergroupLogic) Getlayergroup(req types.GetLayerGroupReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
