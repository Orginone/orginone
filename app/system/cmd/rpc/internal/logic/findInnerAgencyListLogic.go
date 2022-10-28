package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/predicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindInnerAgencyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindInnerAgencyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindInnerAgencyListLogic {
	return &FindInnerAgencyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询部门
func (l *FindInnerAgencyListLogic) FindInnerAgencyList(in *system.FindInnerAgencyListReq) (*system.CommonRpcRes, error) {
	predicates := make([]predicate.AsInnerAgency, 0)
	if len(in.Ids) > 0 {
		predicates = append(predicates, asinneragency.IDIn(in.Ids...))
	} else {
		predicates = append(predicates, asinneragency.TenantCode(in.TenantCode), asinneragency.AgencyNameContains(in.AgencyName), asinneragency.IsDeleted(0))
	}
	query := l.svcCtx.SystemStore.AsInnerAgency.Query().Where(predicates...).WithParent()
	if in.Page != nil {
		count, err := query.Count(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		agencyEntitys, err := query.Order(schema.Asc(asinneragency.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
		return system.PageResult(in.Page, int64(count), agencyEntitys, err)
	}
	return system.JsonResult(query.All(l.ctx))
}
