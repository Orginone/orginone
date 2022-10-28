package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuditUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuditUserLogic {
	return &AuditUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 审核租户
func (l *AuditUserLogic) AuditUser(in *user.AuditUserReq) (*user.CommonRpcRes, error) {
	userUpdate := l.svcCtx.UserStore.AsUser.Update().Where(asuser.IDIn(in.UserIds...), asuser.TenantCode(in.TenantCode))
	personUpdate := l.svcCtx.UserStore.AsPerson.Update().Where(asperson.UserIDIn(in.UserIds...), asperson.TenantCode(in.TenantCode))
	if in.IsPass {
		roleId, err := l.svcCtx.UserStore.GetRoleId(l.ctx, 3)
		if err != nil {
			return &user.CommonRpcRes{}, err
		}
		_, err = userUpdate.SetStatus(2).SetTenantApplyingState(2).AddRoleIDs(roleId).Save(l.ctx)
		if err != nil {
			return &user.CommonRpcRes{}, err
		}
		return user.JsonResult(personUpdate.SetStatus(2).Save(l.ctx))
	} else {
		_, err := userUpdate.SetStatus(0).SetTenantApplyingState(3).Save(l.ctx)
		if err != nil {
			return &user.CommonRpcRes{}, err
		}
		return user.JsonResult(personUpdate.SetIsDeleted(1).Save(l.ctx))
	}
}
