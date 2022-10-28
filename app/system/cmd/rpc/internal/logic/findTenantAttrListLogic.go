package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/astenantattr"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantAttrListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantAttrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantAttrListLogic {
	return &FindTenantAttrListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询租户特性
func (l *FindTenantAttrListLogic) FindTenantAttrList(in *system.FindTenantAttrListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsTenantAttr.Query().Where(astenantattr.AttrNameContains(in.AttrName), astenantattr.IsDeleted(0))
	count, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tenantAttrEntitys, err := query.Order(schema.Asc(astenantattr.FieldAttrName)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(count), tenantAttrEntitys, err)
}
