package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/aslayer"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisbandGroupByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisbandGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisbandGroupByIdLogic {
	return &DisbandGroupByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 解散集团
func (l *DisbandGroupByIdLogic) DisbandGroupById(in *system.DisbandGroupReq) (*system.CommonRpcRes, error) {
	sourceGroupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.SourceGroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.GroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	err = l.DisbandGroupOwner(groupEntity, sourceGroupEntity.ID, in.UserId)
	return system.Result("解散成功.", err)
}

// 解散集团本身
func (l *DisbandGroupByIdLogic) DisbandGroupOwner(groupEntity *schema.AsAllGroup, sourceGroupId int64, userId int64) error {
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return err
	}
	ids := []int64{groupEntity.ID}
	if !tools.IsNilOrEmpty(groupEntity.GroupCode) {
		groupIds, err := tx.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.StatusIn(102, 106),
			asgrouptenantrelations.GroupCodeHasPrefix(groupEntity.GroupCode), asgrouptenantrelations.Type(1)).Select(asgrouptenantrelations.FieldSonID).Int64s(l.ctx)
		if err != nil {
			return err
		}
		_, err = tx.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.StatusIn(102, 106),
			asgrouptenantrelations.GroupCodeHasPrefix(groupEntity.GroupCode)).SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
		ids = append(ids, groupIds...)
	} else {
		_, err = tx.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.StatusIn(102, 106),
			asgrouptenantrelations.ParentID(groupEntity.ID)).SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	_, err = tx.AsAllGroup.Update().Where(asallgroup.IsDeleted(0), asallgroup.IDIn(ids...)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.AsLayer.Update().Where(aslayer.IsDeleted(0), aslayer.GroupID(groupEntity.ID)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if groupEntity.ID == sourceGroupId {
		exist, err := tx.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.TenantCode(groupEntity.TenantCode)).Exist(l.ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
		if !exist {
			attrId, err := l.svcCtx.SystemStore.GetTenantTypeId(l.ctx, 1)
			if err != nil {
				tx.Rollback()
				return err
			}
			groupAttrId, err := l.svcCtx.SystemStore.GetTenantTypeId(l.ctx, 4)
			if err != nil {
				tx.Rollback()
				return err
			}
			_, err = tx.AsTenant.Update().Where(astenant.TenantCode(groupEntity.TenantCode),
				astenant.TenantType(groupAttrId)).SetTenantType(attrId).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return err
			}
			roleId, err := l.svcCtx.SystemStore.GetRoleId(l.ctx, 5)
			if err != nil {
				return err
			}
			_, err = tx.AsUser.Update().Where(asuser.ID(userId)).RemoveRoleIDs(roleId).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Commit()
	return err
}
