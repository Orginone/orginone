package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelJoinTenantRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelJoinTenantRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelJoinTenantRequestLogic {
	return &CancelJoinTenantRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消加入租户申请
func (l *CancelJoinTenantRequestLogic) CancelJoinTenantRequest(in *user.CancelJoinTenantReq) (*user.CommonRpcRes, error) {
	tx, err := l.svcCtx.UserStore.Tx(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	_, err = tx.AsUser.Update().Where(asuser.IsDeleted(0), asuser.PhoneNumber(in.Account), asuser.TenantCodeIn(in.TenantCodes...), asuser.Status(1)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &user.CommonRpcRes{}, err
	}
	_, err = tx.AsPerson.Update().Where(asperson.IsDeleted(0), asperson.PhoneNumber(in.Account), asperson.TenantCodeIn(in.TenantCodes...), asperson.Status(1)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &user.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return &user.CommonRpcRes{}, nil
}
