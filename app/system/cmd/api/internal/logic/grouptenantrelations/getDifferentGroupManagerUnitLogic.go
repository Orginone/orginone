package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDifferentGroupManagerUnitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDifferentGroupManagerUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDifferentGroupManagerUnitLogic {
	return GetDifferentGroupManagerUnitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDifferentGroupManagerUnitLogic) GetDifferentGroupManagerUnit(req types.GetDifferentGroupManagerUnitReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
