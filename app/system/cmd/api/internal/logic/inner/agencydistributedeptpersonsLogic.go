package inner

import (
	"context"
	"net/http"
	"strconv"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencydistributedeptpersonsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencydistributedeptpersonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencydistributedeptpersonsLogic {
	return AgencydistributedeptpersonsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencydistributedeptpersonsLogic) Agencydistributedeptpersons(req types.AgencyDistributeDeptPersonsReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusRequestTimeout, err)
	}
	dept, err := strconv.ParseInt(req.DeptId, 10, 64)
	if err != nil {
		return types.Failed(http.StatusRequestTimeout, err)
	}
	return types.CommonResult(l.svcCtx.SystemRpc.DistributeDeptPersons(l.ctx, &system.DistributeDeptPersonsReq{
		TenantCode:   tenantCode,
		DeptId:       dept,
		Type:         req.Type,
		PersonIdList: tools.StrArrayToNumArray(req.PersonIdList),
	}))
}
