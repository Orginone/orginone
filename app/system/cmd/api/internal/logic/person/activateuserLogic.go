package person

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivateuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivateuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivateuserLogic {
	return ActivateuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivateuserLogic) Activateuser(req []string) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.BoolResult(l.svcCtx.SystemRpc.ActiveUser(l.ctx, &system.ActiveUserReq{
		Phones:     req,
		TenantCode: tenantCode,
	}))

}
