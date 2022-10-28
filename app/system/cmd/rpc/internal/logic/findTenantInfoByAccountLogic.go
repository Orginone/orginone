package logic

import (
	"context"
	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantInfoByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantInfoByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantInfoByAccountLogic {
	return &FindTenantInfoByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据账户查找租户分页
func (l *FindTenantInfoByAccountLogic) FindTenantInfoByAccount(in *system.AccountTenantReq) (*system.CommonRpcRes, error) {
	tenantCodes, err := l.svcCtx.SystemStore.GetTenantByAccountAndCount(l.ctx, in.Account, in.Count)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	query := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.IsDeleted(0), astenant.TenantCodeIn(tenantCodes...), astenant.TenantTypeNEQ(7))
	if !tools.IsNilOrEmpty(in.Page.Filter) {
		query = query.Where(astenant.TenantNameContains(in.Page.Filter))
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tenantEntitys, err := query.Where(astenant.HasUnitWith(asunit.IsDeleted(0))).WithUnit().Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tenantModels := make([]*model.TenantModel, 0)
	for _, tenant := range tenantEntitys {
		item := new(model.TenantModel)
		item.AsTenant = tenant
		item.PhoneNumber = in.Account
		item.RealName = tenant.Edges.Unit.LinkMan
		item.SocialCreCode = tenant.Edges.Unit.SocialCreditCode
		userEntitys, err := l.svcCtx.SystemStore.AsUser.Query().Where(asuser.PhoneNumber(in.Account), asuser.IsDeleted(0), asuser.TenantCode(tenant.TenantCode)).All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if len(userEntitys) > 0 {
			item.IsCreated = userEntitys[0].IsCreated
			item.TenantApplyingState = userEntitys[0].TenantApplyingState
		}
		personEntitys, err := l.svcCtx.SystemStore.AsPerson.Query().Where(asperson.PhoneNumber(in.Account), asperson.IsDeleted(0), asperson.TenantCode(tenant.TenantCode)).All(l.ctx)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		if len(personEntitys) > 0 {
			item.IsMaster = personEntitys[0].IsMaster
		}
		tenantAttr, err := l.svcCtx.SystemStore.AsTenantAttr.Get(l.ctx, tenant.TenantType)
		if err != nil {
			return &system.CommonRpcRes{}, err
		}
		item.TenantTypeName = tenantAttr.AttrName
		tenantModels = append(tenantModels, item)
	}
	return system.PageResult(in.Page, total, tenantModels, err)
}
