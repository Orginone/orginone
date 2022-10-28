package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantCodeControlGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTenantCodeControlGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTenantCodeControlGroupLogic {
	return GetTenantCodeControlGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTenantCodeControlGroupLogic) GetTenantCodeControlGroup(req types.GetTenantCodeControlGroupReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
