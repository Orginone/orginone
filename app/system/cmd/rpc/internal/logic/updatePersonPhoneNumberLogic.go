package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonPhoneNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePersonPhoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonPhoneNumberLogic {
	return &UpdatePersonPhoneNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新人员电话 补充电话
func (l *UpdatePersonPhoneNumberLogic) UpdatePersonPhoneNumber(in *system.UpdatePersonPhoneNumberReq) (*system.CommonRpcRes, error) {
	isExist, err := l.svcCtx.SystemStore.AsPerson.Query().Where(asperson.IDNEQ(in.PersonId),
		asperson.PhoneNumber(in.PhoneNumber), asperson.TenantCode(in.TenantCode), asperson.IsDeleted(0)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if isExist {
		return &system.CommonRpcRes{}, errors.New("该手机号已注册!")
	}
	_, err = l.svcCtx.SystemStore.AsUser.Update().Where(asuser.IsDeleted(0), asuser.HasPersonWith(asperson.IsDeleted(0), asperson.ID(in.PersonId))).SetPhoneNumber(in.PhoneNumber).Save(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(l.svcCtx.SystemStore.AsPerson.UpdateOneID(in.PersonId).SetPhoneNumber(in.PhoneNumber).Save(l.ctx))
}
