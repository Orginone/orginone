package inner

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"

	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyinnertreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyinnertreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyinnertreeLogic {
	return AgencyinnertreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyinnertreeLogic) Agencyinnertree(req types.AgencyInnerTreeReq) (resp *types.CommonResponse, err error) {
	rpcRes, err := l.svcCtx.SystemRpc.FindInnerAgencyTreeByTenantCode(l.ctx, &system.TenantCodeReq{
		TenantCode: req.TenantCode,
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
