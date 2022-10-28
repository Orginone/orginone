package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asgrouptenantrelations"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindGroupListByRelationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindGroupListByRelationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindGroupListByRelationsLogic {
	return &FindGroupListByRelationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有下属集团
func (l *FindGroupListByRelationsLogic) FindGroupListByRelations(in *system.FindGroupListByRelationsReq) (*system.CommonRpcRes, error) {
	allGroupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.GroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if allGroupEntity == nil {
		return &system.CommonRpcRes{}, errors.New("group not found.")
	}
	query := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.ExpiresTimeIsNil(),
		asgrouptenantrelations.StatusIn(102, 106), asgrouptenantrelations.Type(1),
		asgrouptenantrelations.IsDeleted(0), asgrouptenantrelations.GroupCodeHasPrefix(allGroupEntity.GroupCode))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	sonIds, err := query.Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).Select(asgrouptenantrelations.FieldSonID).Int64s(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	groupIds := append(sonIds, in.GroupId)
	data, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.IDIn(groupIds...)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(total+1), data, err)
}
