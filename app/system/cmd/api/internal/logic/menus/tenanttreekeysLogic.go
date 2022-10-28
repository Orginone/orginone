package menus

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenanttreekeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenanttreekeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenanttreekeysLogic {
	return TenanttreekeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenanttreekeysLogic) Tenanttreekeys(req types.TenantTreeKeysReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
