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

type BanuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBanuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) BanuserLogic {
	return BanuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BanuserLogic) Banuser(req []string) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.BoolResult(l.svcCtx.SystemRpc.BanPersonUser(l.ctx, &system.BanPersonUserReq{
		Phones:     req,
		TenantCode: tenantCode,
	}))
}
