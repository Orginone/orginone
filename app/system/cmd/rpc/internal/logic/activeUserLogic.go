package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActiveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveUserLogic {
	return &ActiveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增内部机构并分配人员
func (l *ActiveUserLogic) ActiveUser(in *system.ActiveUserReq) (*system.CommonRpcRes, error) {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsUser.Update().Where(asuser.PhoneNumberIn(in.Phones...), asuser.TenantCode(in.TenantCode), asuser.IsDeleted(0)).SetStatus(2).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsPerson.Update().Where(asperson.PhoneNumberIn(in.Phones...), asperson.TenantCode(in.TenantCode), asperson.IsDeleted(0)).SetStatus(2).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return &system.CommonRpcRes{}, err
}
