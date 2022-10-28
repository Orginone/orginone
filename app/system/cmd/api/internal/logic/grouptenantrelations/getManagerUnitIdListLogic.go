package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManagerUnitIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetManagerUnitIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetManagerUnitIdListLogic {
	return GetManagerUnitIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetManagerUnitIdListLogic) GetManagerUnitIdList(req types.GetManagerUnitIdListReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
