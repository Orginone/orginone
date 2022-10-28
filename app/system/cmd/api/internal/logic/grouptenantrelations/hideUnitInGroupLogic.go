package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type HideUnitInGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHideUnitInGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) HideUnitInGroupLogic {
	return HideUnitInGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HideUnitInGroupLogic) HideUnitInGroup(req types.HideUnitInGroupReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.UpdateGroupIsHide(l.ctx, &system.UpdateGroupIsHideReq{
		Type:          req.Type,
		TenantId:      req.TenantId,
		GroupCode:     req.GroupCode,
		SourceGroupId: req.SourceGroupId,
	}))
}
