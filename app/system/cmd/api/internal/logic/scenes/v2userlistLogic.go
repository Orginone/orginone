package scenes

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2userlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2userlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2userlistLogic {
	return V2userlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2userlistLogic) V2userlist(req types.V2UserListReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindTenantByAccount(l.ctx, &system.ByAccountReq{Account: req.PhoneNumber})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var tenantEntitys []*schema.AsTenant
	err = json.Unmarshal([]byte(rpcres.Json), &tenantEntitys)
	if err != nil {
		return types.Failed(http.StatusNotExtended, err)
	}
	return types.Successful(tenantEntitys)
}
