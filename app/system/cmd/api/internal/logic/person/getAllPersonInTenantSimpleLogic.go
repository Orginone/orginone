package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPersonInTenantSimpleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPersonInTenantSimpleLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllPersonInTenantSimpleLogic {
	return GetAllPersonInTenantSimpleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPersonInTenantSimpleLogic) GetAllPersonInTenantSimple(req types.GetAllPersonInTenantSimpleReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
