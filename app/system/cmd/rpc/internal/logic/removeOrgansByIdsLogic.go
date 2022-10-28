package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveOrgansByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveOrgansByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveOrgansByIdsLogic {
	return &RemoveOrgansByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除单位
func (l *RemoveOrgansByIdsLogic) RemoveOrgansByIds(in *system.V2RemoveByIdsReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsUnit.Update().Where(asunit.IsDeleted(0), asunit.IDIn(in.Ids...)).SetIsDeleted(1).Save(l.ctx))
}
