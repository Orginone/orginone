package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupAllTenantCodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupAllTenantCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetGroupAllTenantCodeListLogic {
	return GetGroupAllTenantCodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupAllTenantCodeListLogic) GetGroupAllTenantCodeList(req types.GetGroupAllTenantCodeListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
