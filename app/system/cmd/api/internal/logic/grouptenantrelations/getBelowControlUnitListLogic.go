package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBelowControlUnitListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBelowControlUnitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBelowControlUnitListLogic {
	return GetBelowControlUnitListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBelowControlUnitListLogic) GetBelowControlUnitList(req types.GetBelowControlUnitListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
