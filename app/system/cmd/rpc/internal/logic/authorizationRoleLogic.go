package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asmenu"
	"orginone/common/schema/asrole"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthorizationRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizationRoleLogic {
	return &AuthorizationRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设置角色权限
func (l *AuthorizationRoleLogic) AuthorizationRole(in *system.AuthorizationRoleReq) (*system.CommonRpcRes, error) {
	ids, err := l.svcCtx.SystemStore.AsMenu.Find().Where(asmenu.IDIn(in.MenuIds...), asmenu.HasParentx()).QueryParentx().Select(asmenu.FieldID).Int64s(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	in.MenuIds = append(in.MenuIds, ids...)
	linq.From(in.MenuIds).Distinct().ToSlice(&in.MenuIds)
	return system.JsonResult(l.svcCtx.SystemStore.AsRole.Update().Where(asrole.ID(in.RoleId)).ClearMenus().AddMenuIDs(in.MenuIds...).Save(l.ctx))
}
