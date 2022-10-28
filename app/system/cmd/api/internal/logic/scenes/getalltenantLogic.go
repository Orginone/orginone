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

type GetalltenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetalltenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetalltenantLogic {
	return GetalltenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetalltenantLogic) Getalltenant(req types.GetAllTenantReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.SearchTeantByName(context.Background(), &system.SearchTeantReq{TenantName: req.TenantName})
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	} else {
		var tenantEntitys []*schema.AsTenant
		err = json.Unmarshal([]byte(rpcres.Json), &tenantEntitys)
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		} else {
			return types.Successful(tenantEntitys)
		}
	}
}
