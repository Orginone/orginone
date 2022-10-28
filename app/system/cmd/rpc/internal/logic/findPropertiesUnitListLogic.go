package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/aspropertiesdistribution"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPropertiesUnitListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPropertiesUnitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPropertiesUnitListLogic {
	return &FindPropertiesUnitListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取性质里的租户列表分页
func (l *FindPropertiesUnitListLogic) FindPropertiesUnitList(in *system.FindPropertiesUnitListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsPropertiesDistribution.Query().
		Where(
			aspropertiesdistribution.PropertiesID(in.PropertiesId),
			aspropertiesdistribution.IsDeleted(0),
		).
		QueryTenant().Where(astenant.IsDeleted(0)).
		QueryUnit().Where(asunit.IsDeleted(0))

	if in.Page != nil {
		total, err := query.Count(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		unitEntitys, err := query.
			Order(schema.Desc(asunit.FieldEnableTime)).
			Offset(int(in.Page.Offset)).
			Limit(int(in.Page.Limit)).
			All(l.ctx)
		return system.PageResult(in.Page, total, unitEntitys, err)
	} else {
		return system.JsonResult(query.All(l.ctx))
	}
}
