package person

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchdeptpersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchdeptpersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchdeptpersonLogic {
	return SearchdeptpersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchdeptpersonLogic) Searchdeptperson(req types.SearchDeptPersonReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	data := make(map[string]interface{}, 0)
	var persons []schema.AsPerson
	rpcRes, err := l.svcCtx.SystemRpc.SearchPersonByName(l.ctx, &system.SerarchDeptPersonReq{Filter: req.Name, TenantCode: tenantCode})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	err = json.Unmarshal([]byte(rpcRes.Json), &persons)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	data["personList"] = persons
	var agencys []schema.AsInnerAgency
	rpcRes, err = l.svcCtx.SystemRpc.SearchDeptByName(l.ctx, &system.SerarchDeptPersonReq{Filter: req.Name, TenantCode: tenantCode})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	err = json.Unmarshal([]byte(rpcRes.Json), &agencys)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	data["deptList"] = agencys
	return types.JsonResult(data, nil)
}
