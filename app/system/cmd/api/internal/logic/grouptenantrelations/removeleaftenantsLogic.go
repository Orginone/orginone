package grouptenantrelations

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveleaftenantsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveleaftenantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveleaftenantsLogic {
	return RemoveleaftenantsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveleaftenantsLogic) Removeleaftenants(req types.RemoveLeafTenantsReq) (resp *types.CommonResponse, err error) {
	gId, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	return types.BoolResult(l.svcCtx.SystemRpc.RemoveLeafTenants(l.ctx, &system.RemoveLeafTenantsReq{
		TenantIds: tools.ArrayStrToInt64(strings.Split(req.TenantIds, ",")),
		GroupId:   gId,
	}))
}
