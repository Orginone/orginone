package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/schema/asuser"
	"orginone/common/schema/asworkingdata"
	"orginone/common/schema/predicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAgencyRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgencyRemoveLogic {
	return &AgencyRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除部门
func (l *AgencyRemoveLogic) AgencyRemove(in *system.AgencyRemoveReq) (*system.CommonRpcRes, error) {
	// return system.JsonResult(l.svcCtx.SystemStore.AsInnerAgency.Remove().
	// 	Where(
	// 		asinneragency.IDIn(in.Ids...),
	// 		asinneragency.Not(asinneragency.HasChildrens()),
	// 		asinneragency.Not(asinneragency.HasRoleDistribs()),
	// 		asinneragency.Not(asinneragency.HasUsersWith(asuser.HasWorkingDatasWith(asworkingdata.Type(2)))),
	// 	).
	// 	ClearPersons().
	// 	RemoveRoleDistribs().
	// 	Save(l.ctx))
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	pre := []predicate.AsInnerAgency{
		asinneragency.IDIn(in.Ids...),
		asinneragency.Not(asinneragency.HasChildrens()),
		asinneragency.Not(asinneragency.HasRoleDistribs()),
		asinneragency.Not(asinneragency.HasUsersWith(asuser.HasWorkingDatasWith(asworkingdata.Type(2))))}
	ids, err := tx.AsInnerAgency.Query().Where(pre...).Select(asinneragency.FieldID).Int64s(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsInnerAgency.Remove().Where(pre...).ClearPersons().Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketRoleDistribution.Remove().Where(asmarketroledistribution.AgencyIDIn(ids...)).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(true, tx.Commit())
}
