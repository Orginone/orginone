package grouptenantrelations

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatefathertenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatefathertenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatefathertenantLogic {
	return UpdatefathertenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatefathertenantLogic) Updatefathertenant(req types.UpdateFatherTenantReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.UpdateFatherTenant(l.ctx, &system.UpdateFatherTenantReq{
		TenantIds: tools.ArrayStrToInt64(strings.Split(req.TenantIds, ",")),
		GroupId:   req.GroupId,
	}))
}
