package tenanticon

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

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListReq3) (resp *types.CommonResponse, err error) {
	// 查询人员相关
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	rpcres, err := l.svcCtx.SystemRpc.FindTenantIconList(l.ctx, &system.FindTenantIconListReq{TenantCode: tenantCode, TenantIcon: req.TenantIcon})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	icons := make([]*schema.AsTenantIcon, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &icons)
	return types.JsonResult(icons, err)
}
