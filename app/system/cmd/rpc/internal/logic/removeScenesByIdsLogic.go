package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveScenesByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveScenesByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveScenesByIdsLogic {
	return &RemoveScenesByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除租户
func (l *RemoveScenesByIdsLogic) RemoveScenesByIds(in *system.RemoveScenesByIdsReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsTenant.Update().Where(astenant.IsDeleted(0), astenant.IDIn(in.Ids...)).SetIsDeleted(1).Save(l.ctx))
}
