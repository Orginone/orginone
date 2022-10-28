package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPersonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPersonLogic {
	return &AddPersonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增组织人员
func (l *AddPersonLogic) AddPerson(in *system.AddPersonReq) (*system.CommonRpcRes, error) {
	isExist, err := l.svcCtx.SystemStore.AsPerson.Query().
		Where(
			asperson.IsDeleted(0),
			asperson.TenantCode(in.TenantCode),
			asperson.Or(
				asperson.UserCode(in.UserCode),
				asperson.PhoneNumber(in.Phone),
			),
		).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if isExist {
		return &system.CommonRpcRes{}, errors.New("该编号或电话号码已存在!")
	}
	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if l.svcCtx.SystemStore.CreatePerson(l.ctx, tx, in.TenantCode, in.Name, in.UserCode, in.Phone, []int64{in.Role}, in.Depart) != nil {
		tx.Rollback()
	}
	return system.Result("true", tx.Commit())
}
