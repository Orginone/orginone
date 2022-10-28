package inner

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetDeptCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetDeptCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetDeptCodeLogic {
	return AgencygetDeptCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetDeptCodeLogic) AgencygetDeptCode(req types.AgencyGetDeptCodeReq) (resp *types.CommonResponse, err error) {
	rpcRes, err := l.svcCtx.SystemRpc.FindAgencyDeptCodeByTenantCode(l.ctx, &system.TenantCodeReq{
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(rpcRes.Json)
}
