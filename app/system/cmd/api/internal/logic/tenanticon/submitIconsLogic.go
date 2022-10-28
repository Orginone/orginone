package tenanticon

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitIconsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitIconsLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitIconsLogic {
	return SubmitIconsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitIconsLogic) SubmitIcons(req types.SubmitIconsReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.ImportTenantIcon(l.ctx, &system.ImportTenantIconReq{
		TenantCode: req.TenantCode,
		Urls:       strings.Split(req.Urls, ","),
	}))
}
