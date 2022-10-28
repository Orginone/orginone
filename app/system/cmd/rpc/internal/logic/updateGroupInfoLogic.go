package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/tools"
	"orginone/common/tools/date"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupInfoLogic {
	return &UpdateGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新集团信息
func (l *UpdateGroupInfoLogic) UpdateGroupInfo(in *system.UpdateGroupInfoReq) (*system.CommonRpcRes, error) {
	groupEntity, err := l.checkGroupInfoData(in)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	updater := groupEntity.Update().SetGroupName(in.GroupName)
	if !tools.IsNilOrEmpty(in.GroupDescription) {
		updater = updater.SetGroupDescription(in.GroupDescription)
	}
	if !tools.IsNilOrEmpty(in.SocialCreditCode) && tools.IsSocialCreditCode(in.SocialCreditCode) {
		updater = updater.SetSocialCreditCode(in.SocialCreditCode)
	}
	if in.TenantCode != groupEntity.TenantCode {
		updater = updater.SetTenantCode(in.TenantCode)
	}
	groupEntity, err = updater.Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	if in.ParentGroupId > 0 && groupEntity.Type == 2 {
		exist, err := tx.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.ExpiresTimeIsNil(),
			asgrouptenantrelations.ParentID(in.ParentGroupId), asgrouptenantrelations.SonID(in.SonGroupId), asgrouptenantrelations.Type(1)).Exist(l.ctx)
		if err != nil {
			tx.Rollback()
			return &system.CommonRpcRes{}, err
		}
		if !exist {
			parentGroupEntity, err := tx.AsAllGroup.Get(l.ctx, in.ParentGroupId)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			sourceGroupId, err := l.svcCtx.SystemStore.GetTopGroupId(l.ctx, groupEntity.ID)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			sourceGroupEntity, err := tx.AsAllGroup.Get(l.ctx, sourceGroupId)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			_, err = tx.AsGroupTenantRelations.Update().Where(asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.SonID(groupEntity.ID),
				asgrouptenantrelations.GroupCodeHasPrefix(sourceGroupEntity.GroupCode)).SetExpiresTime(date.Now()).Save(l.ctx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
			err = l.svcCtx.SystemStore.GroupJoinGroup(l.ctx, groupEntity, parentGroupEntity, tx)
			if err != nil {
				tx.Rollback()
				return &system.CommonRpcRes{}, err
			}
		}
	}
	err = tx.Commit()
	return system.Result("true", err)
}

//检验数据
func (l *UpdateGroupInfoLogic) checkGroupInfoData(in *system.UpdateGroupInfoReq) (*schema.AsAllGroup, error) {
	groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.SonGroupId)
	if err != nil {
		return nil, err
	}
	if groupEntity.GroupName != in.GroupName {
		exist, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.GroupName(in.GroupName)).Exist(l.ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, errors.New("name is used, please change it.")
		}
	}
	return groupEntity, err
}
