package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUnitByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUnitByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUnitByIdsLogic {
	return &FindUnitByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有单位列表（包含子集团）根据集团id
func (l *FindUnitByIdsLogic) FindUnitByIds(in *system.FindUnitByIdsReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.IDIn(in.UnitIdList...)).All(l.ctx))
}
