package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetunitbycodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetunitbycodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetunitbycodeLogic {
	return GetunitbycodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetunitbycodeLogic) Getunitbycode(req types.GetUnitByCodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
