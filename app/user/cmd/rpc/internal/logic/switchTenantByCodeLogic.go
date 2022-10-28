package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchTenantByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSwitchTenantByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchTenantByCodeLogic {
	return &SwitchTenantByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 切换主单位
func (l *SwitchTenantByCodeLogic) SwitchTenantByCode(in *user.SwitchTenantReq) (*user.CommonRpcRes, error) {
	//开启事务
	tx, err := l.svcCtx.UserStore.Tx(l.ctx)
	if err != nil {
		return user.Result("false", err)
	}
	_, err = tx.AsUser.Update().Where(asuser.PhoneNumber(in.Account)).SetIsMaster(0).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return user.Result("false", err)
	}
	_, err = tx.AsUser.Update().Where(asuser.PhoneNumber(in.Account), asuser.TenantCode(in.TenantCode)).SetIsMaster(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return user.Result("false", err)
	}
	err = tx.Commit()
	if err != nil {
		return user.Result("false", err)
	}
	return user.JsonResult(l.svcCtx.UserStore.AsUser.Query().Where(asuser.IsMaster(1), asuser.IsDeleted(0), asuser.PhoneNumber(in.Account)).WithRoles().First(l.ctx))
}
