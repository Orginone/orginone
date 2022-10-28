package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpGroupManagerUnitIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpGroupManagerUnitIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUpGroupManagerUnitIdLogic {
	return GetUpGroupManagerUnitIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUpGroupManagerUnitIdLogic) GetUpGroupManagerUnitId(req types.GetUpGroupManagerUnitIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
