package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type BanPersonUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBanPersonUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BanPersonUserLogic {
	return &BanPersonUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 人员停用
func (l *BanPersonUserLogic) BanPersonUser(in *system.BanPersonUserReq) (*system.CommonRpcRes, error) {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsUser.Update().SetStatus(3).Where(
		asuser.IsDeleted(0),
		asuser.PhoneNumberIn(in.Phones...),
		asuser.TenantCode(in.TenantCode),
	).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsPerson.Update().SetStatus(3).Where(
		asperson.IsDeleted(0),
		asperson.PhoneNumberIn(in.Phones...),
		asperson.TenantCode(in.TenantCode),
	).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return &system.CommonRpcRes{}, err
}
