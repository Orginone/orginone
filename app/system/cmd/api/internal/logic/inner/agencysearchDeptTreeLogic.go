package inner

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencysearchDeptTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencysearchDeptTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencysearchDeptTreeLogic {
	return AgencysearchDeptTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencysearchDeptTreeLogic) AgencysearchDeptTree(req types.AgencySearchDeptTreeReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	rpcRes, err := l.svcCtx.SystemRpc.FindInnerAgencyTreeByTenantCode(l.ctx, &system.TenantCodeReq{
		TenantCode: tenantCode,
		Filter:     req.Name,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var innerAgencyTrees []*model.InnerAgencyTree
	err = json.Unmarshal([]byte(rpcRes.Json), &innerAgencyTrees)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(innerAgencyTrees)
}
