package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindGroupLayerByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindGroupLayerByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindGroupLayerByIdLogic {
	return &FindGroupLayerByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询集团层级根据集团id
func (l *FindGroupLayerByIdLogic) FindGroupLayerById(in *system.GroupLayerByIdReq) (*system.CommonRpcRes, error) {
	groupInfos := make([]*model.GroupUnitInfo, 0)
	sourceGroupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, in.SourceGroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if in.WithUnit {
		//单位类型节点
		unitNodeInfos, err := l.svcCtx.SystemStore.FindUnitNodeInfoById(l.ctx, sourceGroupEntity.TenantCode, in.GroupId)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if len(unitNodeInfos) > 0 {
			groupInfos = append(groupInfos, unitNodeInfos...)
		}
	}
	//集团类型节点
	groupNodeInfos, err := l.svcCtx.SystemStore.FindGroupNodeInfoById(l.ctx, sourceGroupEntity.TenantCode, in.GroupId)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if len(groupNodeInfos) > 0 {
		groupInfos = append(groupInfos, groupNodeInfos...)
	}
	return system.JsonResult(groupInfos, err)
}
