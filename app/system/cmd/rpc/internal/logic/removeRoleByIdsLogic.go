package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asrole"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRoleByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveRoleByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRoleByIdsLogic {
	return &RemoveRoleByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *RemoveRoleByIdsLogic) RemoveRoleByIds(in *system.RemoveRoleByIdsReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsRole.Update().Where(asrole.IDIn(in.RoleIds...)).
		ClearMenus().ClearJobs().ClearUsers().ClearAttrRoles().SetIsDeleted(1).Save(l.ctx))
}
