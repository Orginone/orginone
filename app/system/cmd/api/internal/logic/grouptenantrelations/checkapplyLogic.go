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

type CheckapplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckapplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckapplyLogic {
	return CheckapplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckapplyLogic) Checkapply(req types.CheckApplyReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.TenantsJoinGroup(l.ctx, &system.TenantsJoinGroupReq{
		GroupId:   req.GroupId,
		Flag:      req.Flag,
		TenantIds: tools.ArrayStrToInt64(strings.Split(req.TenantIds, ",")),
	}))
}
