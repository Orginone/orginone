package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPersonInTenantALogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPersonInTenantALogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllPersonInTenantALogic {
	return GetAllPersonInTenantALogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPersonInTenantALogic) GetAllPersonInTenantA(req types.GetAllPersonInTenantReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
