package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupTreeLogic {
	return &GetGroupTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取组织树
func (l *GetGroupTreeLogic) GetGroupTree(in *system.GroupTreeReq) (*system.CommonRpcRes, error) {
	groupInfos := make([]*model.GroupUnitInfo, 0)
	groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.SourceGroupId)
	if err != nil {
		return system.JsonResult(groupInfos, err)
	}
	tenantEntity, err := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantCode(groupEntity.TenantCode),
		astenant.IsDeleted(0), astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().First(l.ctx)
	if err != nil {
		return system.JsonResult(groupInfos, err)
	}
	nodes, err := l.GetGroupTreeNodes(in.SourceGroupId, in.GroupId)
	if err != nil {
		return system.JsonResult(groupInfos, err)
	}
	groupInfos = append(groupInfos, model.NewGroupUnitInfo(1, len(nodes) > 0).SetChildrens(nodes...).
		SetParentId(-1).SetTenant(tenantEntity).SetGroup(groupEntity))
	return system.JsonResult(groupInfos, err)
}

func (l *GetGroupTreeLogic) GetGroupTreeNodes(groupId int64, excludeId int64) ([]*model.GroupUnitInfo, error) {
	groupInfos := make([]*model.GroupUnitInfo, 0)
	groupNodeInfos, err := l.svcCtx.SystemStore.FindGroupNodeInfoById(l.ctx, "", groupId)
	if err != nil {
		return groupNodeInfos, err
	}
	for _, node := range groupNodeInfos {
		if node.Id == excludeId {
			continue
		}
		if node.Below {
			node.Children, err = l.GetGroupTreeNodes(node.Id, excludeId)
		}
		groupInfos = append(groupInfos, node)
	}
	return groupInfos, err
}
