package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnitLayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnitLayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUnitLayerLogic {
	return GetUnitLayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnitLayerLogic) GetUnitLayer(req types.GetUnitLayerReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
