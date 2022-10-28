package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2getUnitCodeMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2getUnitCodeMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2getUnitCodeMapLogic {
	return V2getUnitCodeMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2getUnitCodeMapLogic) V2getUnitCodeMap(req types.V2GetUnitCodeMapReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
