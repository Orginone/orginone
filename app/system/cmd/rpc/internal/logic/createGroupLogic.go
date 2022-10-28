package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建集团
func (l *CreateGroupLogic) CreateGroup(in *system.CreateGroupReq) (*system.CommonRpcRes, error) {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return nil, err
	}
	err = l.existValidate(in, tx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	groupEntity, err := l.svcCtx.SystemStore.CreateAndUpdateTenant(l.ctx, in, tx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	if in.GroupId > 0 {
		parentGroup, err := tx.AsAllGroup.Get(l.ctx, in.GroupId)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, err
		}
		err = l.svcCtx.SystemStore.GroupJoinGroup(l.ctx, groupEntity, parentGroup, tx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	err = l.svcCtx.SystemStore.CreateGroupTenantRelation(l.ctx, groupEntity.ID, []int64{}, 102)
	return &system.CommonRpcRes{Json: "true"}, nil
}

// 重复验证
func (l *CreateGroupLogic) existValidate(in *system.CreateGroupReq, tx *schema.Tx) error {
	exist, err := tx.AsAllGroup.Query().Where(asallgroup.GroupName(in.GroupName), asallgroup.Type(1), asallgroup.IsDeleted(0)).Exist(l.ctx)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("group name is exist, please change it.")
	}
	if !tools.IsNilOrEmpty(in.GroupCode) {
		exist, err = tx.AsAllGroup.Query().Where(asallgroup.Type(1), asallgroup.IsDeleted(0),
			asallgroup.GroupCodeContains(in.GroupCode)).Exist(l.ctx)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("group code is exist, please change it.")
		}
	}
	return nil
}
