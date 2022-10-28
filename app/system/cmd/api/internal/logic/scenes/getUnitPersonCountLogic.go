package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnitPersonCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnitPersonCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUnitPersonCountLogic {
	return GetUnitPersonCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnitPersonCountLogic) GetUnitPersonCount(req types.GetUnitPersonCountReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
