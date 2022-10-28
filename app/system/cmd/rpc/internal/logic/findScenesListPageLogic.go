package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/astenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindScenesListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindScenesListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindScenesListPageLogic {
	return &FindScenesListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询租户
func (l *FindScenesListPageLogic) FindScenesListPage(in *system.FindScenesListPageReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsTenant.Query().Where(astenant.TenantCodeContains(in.TenantCode),
		astenant.TenantNameContains(in.TenantName))
	if in.Id != -1 {
		query.Where(astenant.ID(in.Id))
	}
	if in.IsDeleted != -1 {
		query.Where(astenant.IsDeleted(in.IsDeleted))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	tenantEntitys, err := query.Order(schema.Asc(astenant.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(count), tenantEntitys, err)
}