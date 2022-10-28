package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asrole"
	"orginone/common/schema/astenantattrrole"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleByAttrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRoleByAttrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleByAttrLogic {
	return &ListRoleByAttrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取租户组角色
func (l *ListRoleByAttrLogic) ListRoleByAttr(in *system.ListRoleByAttrReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsRole.Query().Where(asrole.HasAttrRolesWith(astenantattrrole.AttrID(in.AttrId))).All(l.ctx))
}
