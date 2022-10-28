package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAllGroupByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAllGroupByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAllGroupByIdsLogic {
	return &RemoveAllGroupByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除集团
func (l *RemoveAllGroupByIdsLogic) RemoveAllGroupByIds(in *system.RemoveAllGroupByIdsReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsAllGroup.Update().Where(asallgroup.IsDeleted(0), asallgroup.IDIn(in.Ids...)).SetIsDeleted(1).Save(l.ctx))
}
