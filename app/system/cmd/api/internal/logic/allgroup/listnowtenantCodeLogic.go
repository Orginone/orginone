package allgroup

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

type ListnowtenantCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListnowtenantCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListnowtenantCodeLogic {
	return ListnowtenantCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListnowtenantCodeLogic) ListnowtenantCode(req types.ListNowTenantCodeReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	rpcres, err := l.svcCtx.SystemRpc.FindAllGroupByCodeAndType(l.ctx, &system.SearchAllGroupReq{TenantCode: tenantCode, GroupType: req.Type})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var groupEntitys []*schema.AsAllGroup
	err = json.Unmarshal([]byte(rpcres.Json), &groupEntitys)
	if err != nil {
		return types.Failed(http.StatusNotExtended, err)
	}
	return types.Successful(groupEntitys)
}
