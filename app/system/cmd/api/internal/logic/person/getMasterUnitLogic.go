package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMasterUnitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMasterUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMasterUnitLogic {
	return GetMasterUnitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMasterUnitLogic) GetMasterUnit(req types.GetMasterUnitReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
