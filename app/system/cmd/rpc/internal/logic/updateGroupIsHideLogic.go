package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/astenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupIsHideLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupIsHideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupIsHideLogic {
	return &UpdateGroupIsHideLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 隐藏、显示树形单位
func (l *UpdateGroupIsHideLogic) UpdateGroupIsHide(in *system.UpdateGroupIsHideReq) (*system.CommonRpcRes, error) {
	isExist, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.ID(in.SourceGroupId), asallgroup.Type(1)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if !isExist {
		return &system.CommonRpcRes{}, errors.New("传参错误!找不到根集团!")
	}
	return system.JsonResult(l.svcCtx.SystemStore.AsGroupTenantRelations.Update().
		Where(
			asgrouptenantrelations.GroupCodeContains(in.GroupCode),
			asgrouptenantrelations.IsDeleted(0),
			asgrouptenantrelations.Type(2),
			asgrouptenantrelations.ExpiresTimeIsNil(),
			asgrouptenantrelations.StatusIn(102, 106),
			asgrouptenantrelations.HasTenantWith(astenant.ID(in.TenantId)),
		).SetIsHide(in.Type).Save(l.ctx))
}
