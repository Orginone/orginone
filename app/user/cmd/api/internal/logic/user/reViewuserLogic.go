package user

import (
	"context"
	"strings"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common/rpcs/user"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReViewuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReViewuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReViewuserLogic {
	return ReViewuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReViewuserLogic) ReViewuser(req types.ReViewUserReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.UserRpc.AuditUser(l.ctx, &user.AuditUserReq{
		UserIds:    tools.StrArrayToNumArray(strings.Split(req.UserIds, ",")),
		TenantCode: req.TenantCode,
		IsPass:     req.IsPass == 1,
	}))
}
