package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePersonByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemovePersonByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePersonByIdsLogic {
	return &RemovePersonByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除人员
func (l *RemovePersonByIdsLogic) RemovePersonByIds(in *system.RemovePersonByIdsReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsPerson.Update().Where(asperson.IsDeleted(0), asperson.IDIn(in.Ids...)).SetIsDeleted(1).Save(l.ctx))
}
