package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/astenant"
	"orginone/common/tools/date"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupRelationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupRelationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupRelationsLogic {
	return &AddGroupRelationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部门导入
func (l *AddGroupRelationsLogic) AddGroupRelations(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	l.ctx = context.Background()
	execlData := make([]*model.GroupRelationData, 0)
	err := json.Unmarshal([]byte(in.Json), &execlData)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	groupNames := make([]string, 0)
	for _, v := range execlData {
		groupNames = append(groupNames, v.GroupName, v.ParentGroupName)
	}
	groupEntitys, err := l.svcCtx.SystemStore.AsAllGroup.Query().
		Where(
			asallgroup.IsDeleted(0),
			asallgroup.TenantCode(execlData[0].TenantCode),
			asallgroup.GroupNameIn(groupNames...),
		).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	manageTenant, err := tx.AsTenant.Query().Where(astenant.TenantCode(execlData[0].TenantCode), astenant.IsDeleted(0)).WithAllGroups().First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	serial := int64(len(manageTenant.Edges.AllGroups)) + 1
	for _, v := range execlData {
		var groupEntity, parentGroup *schema.AsAllGroup
		for _, group := range groupEntitys {
			if v.GroupName == group.GroupName {
				groupEntity = group
				if parentGroup != nil {
					break
				}
			}
			if v.ParentGroupName == group.GroupName {
				parentGroup = group
				if groupEntity != nil {
					break
				}
			}
		}
		if parentGroup != nil {
			if groupEntity != nil {
				hasChild, err := tx.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.Type(1), asgrouptenantrelations.ParentID(groupEntity.ID)).Exist(l.ctx)
				if err != nil {
					tx.Rollback()
					return &system.CommonRpcRes{}, err
				}
				if hasChild {
					continue
				}
				//集团加入集团
				err = groupJoinGroup(false, context.Background(), groupEntity, parentGroup, tx)
				if err != nil {
					tx.Rollback()
					return &system.CommonRpcRes{}, err
				}
			} else { //创建新集团
				groupEntity, err = l.svcCtx.SystemStore.CreateAndUpdateTenant(l.ctx, &system.CreateGroupReq{
					GroupName:  v.GroupName,
					TenantCode: v.TenantCode,
					UserId:     v.UserId,
				}, tx)
				if err != nil {
					tx.Rollback()
					return &system.CommonRpcRes{}, err
				}
				//集团加入集团
				err = groupJoinGroup(true, context.Background(), groupEntity, parentGroup, tx)
				if err != nil {
					tx.Rollback()
					return &system.CommonRpcRes{}, err
				}
				groupEntitys = append(groupEntitys, groupEntity)
			}
			groupCode := parentGroup.GroupCode + fmt.Sprintf("%03d", serial)
			_, err = tx.AsGroupTenantRelations.Create().SetParentID(groupEntity.ID).SetSonID(manageTenant.ID).
				SetSerial(serial).SetSort(serial).SetGroupCode(groupCode).SetStatus(106).SetType(2).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			serial++
		}
	}
	return system.JsonResult(true, tx.Commit())
}

// 集团加入集团
func groupJoinGroup(isNewGroup bool, ctx context.Context, groupEntity *schema.AsAllGroup, parentGroup *schema.AsAllGroup, tx *schema.Tx) error {
	serial, err := tx.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.ParentID(parentGroup.ID)).Count(ctx)
	if err != nil {
		return err
	}
	serial = serial + 1
	groupCode := parentGroup.GroupCode + fmt.Sprintf("%03d", serial)
	if !isNewGroup {
		_, err := tx.AsGroupTenantRelations.Update().Where(
			asgrouptenantrelations.IsDeleted(0),
			asgrouptenantrelations.Or(
				asgrouptenantrelations.ParentID(groupEntity.ID),
				asgrouptenantrelations.SonID(groupEntity.ID))).
			SetIsDeleted(1).
			SetExpiresTime(date.DateTime(time.Now())).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	_, err = tx.AsGroupTenantRelations.Create().SetParentID(parentGroup.ID).SetSonID(groupEntity.ID).SetType(1).
		SetStatus(102).SetGroupCode(groupCode).SetSerial(serial).SetSort(serial).Save(ctx)
	if err != nil {
		return err
	}
	_, err = tx.AsAllGroup.UpdateOne(groupEntity).SetGroupCode(groupCode).Save(ctx)
	groupEntity.GroupCode = groupCode
	return err
}
