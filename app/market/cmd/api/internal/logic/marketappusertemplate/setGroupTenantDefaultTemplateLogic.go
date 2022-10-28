package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetGroupTenantDefaultTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetGroupTenantDefaultTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetGroupTenantDefaultTemplateLogic {
	return SetGroupTenantDefaultTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetGroupTenantDefaultTemplateLogic) SetGroupTenantDefaultTemplate(req types.SetGroupTenantDefaultTemplateReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
