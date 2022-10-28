package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asinneragency"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchDeptByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchDeptByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchDeptByNameLogic {
	return &SearchDeptByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分配名字查询部门
func (l *SearchDeptByNameLogic) SearchDeptByName(in *system.SerarchDeptPersonReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.AgencyNameContains(in.Filter), asinneragency.IsDeleted(0), asinneragency.TenantCode(in.TenantCode)).All(l.ctx))
}
