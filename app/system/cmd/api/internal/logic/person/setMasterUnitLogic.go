package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMasterUnitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMasterUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetMasterUnitLogic {
	return SetMasterUnitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMasterUnitLogic) SetMasterUnit(req types.SetMasterUnitReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.SetMasterUnit(l.ctx, &system.SetMasterUnitReq{
		Account:    req.PhoneNumber,
		TenantCode: req.TenantCode,
	}))
}
