package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asrole"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListByUserIdLogic {
	return &FindRoleListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页角色
func (l *FindRoleListByUserIdLogic) FindRoleListByUserId(in *system.FindRoleListByUserIdReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsUser.Query().
		Where(asuser.ID(in.UserId), asuser.IsDeleted(0)).
		QueryRoles().Where(asrole.IsDeleted(0))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	roleEntitys, err := query.Order(schema.Asc(asrole.FieldSort)).Limit(int(in.Page.Limit)).Offset(int(in.Page.Offset)).All(l.ctx)
	return system.PageResult(in.Page, total, roleEntitys, err)
}
