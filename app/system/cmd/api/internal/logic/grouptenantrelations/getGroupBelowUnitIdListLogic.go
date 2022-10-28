package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupBelowUnitIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupBelowUnitIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetGroupBelowUnitIdListLogic {
	return GetGroupBelowUnitIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupBelowUnitIdListLogic) GetGroupBelowUnitIdList(req types.GetGroupBelowUnitIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
