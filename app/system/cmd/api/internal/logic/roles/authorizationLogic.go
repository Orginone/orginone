package roles

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthorizationLogic {
	return AuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorizationLogic) Authorization(req types.AuthorizationReq) (resp *types.CommonResponse, err error) {
	return types.JsonResult(l.svcCtx.SystemRpc.AuthorizationRole(l.ctx, &system.AuthorizationRoleReq{
		RoleId:  req.RoleId,
		MenuIds: tools.ArrayStrToInt64(strings.Split(req.MenuIds, ",")),
	}))
}
