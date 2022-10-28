package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asrole"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListPageLogic {
	return &FindRoleListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询角色
func (l *FindRoleListPageLogic) FindRoleListPage(in *system.FindRoleListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsRole.Query().Where(asrole.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.RoleAlias) {
		query = query.Where(asrole.RoleAliasContains(in.RoleAlias))
	}
	if !tools.IsNilOrEmpty(in.RoleName) {
		query = query.Where(asrole.RoleNameContains(in.RoleName))
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	roleEntitys, err := query.Order(schema.Asc(asrole.FieldSort)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, total, roleEntitys, err)
}
