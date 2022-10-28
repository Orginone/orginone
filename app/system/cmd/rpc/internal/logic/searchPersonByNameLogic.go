package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPersonByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPersonByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPersonByNameLogic {
	return &SearchPersonByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据名字查询人员
func (l *SearchPersonByNameLogic) SearchPersonByName(in *system.SerarchDeptPersonReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsPerson.Query().Where(asperson.RealNameContains(in.Filter), asperson.IsDeleted(0), asperson.TenantCode(in.TenantCode)).All(l.ctx))
}
