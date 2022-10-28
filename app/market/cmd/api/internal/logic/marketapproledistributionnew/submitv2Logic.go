package marketapproledistributionnew

import (
	"context"
	"net/http"
	"strings"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type Submitv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Submitv2Logic {
	return Submitv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Submitv2Logic) Submitv2(req []*types.SubmitV2Req) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.CommonResult(l.svcCtx.MarketRpc.AppRoleDistribe(l.ctx, &market.AppRoleDistribeRpcReq{
		TenantCode: tenantCode,
		Data:       getAppRoleRequest(req),
	}))
}

//获取请求数据并解析
func getAppRoleRequest(req []*types.SubmitV2Req) []*market.AppRoleDistribeReq {
	results := make([]*market.AppRoleDistribeReq, 0)
	for _, item := range req {
		m := new(market.AppRoleDistribeReq)
		m.RoleId = item.RoleId
		m.JobIds = tools.ArrayStrToInt64(strings.Split(item.JobIds, ","))
		m.UserIds = tools.ArrayStrToInt64(strings.Split(item.UserIds, ","))
		m.AgencyIds = tools.ArrayStrToInt64(strings.Split(item.AgencyIds, ","))
		results = append(results, m)
	}
	return results
}
