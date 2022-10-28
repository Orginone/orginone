package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllControlUnitListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllControlUnitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllControlUnitListLogic {
	return GetAllControlUnitListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllControlUnitListLogic) GetAllControlUnitList(req types.GetAllControlUnitListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
