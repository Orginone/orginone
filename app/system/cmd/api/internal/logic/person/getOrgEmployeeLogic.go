package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrgEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrgEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrgEmployeeLogic {
	return GetOrgEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrgEmployeeLogic) GetOrgEmployee(req types.GetOrgEmployeeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
