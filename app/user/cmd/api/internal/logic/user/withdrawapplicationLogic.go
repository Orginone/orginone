package user

import (
	"context"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawapplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawapplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) WithdrawapplicationLogic {
	return WithdrawapplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WithdrawapplicationLogic) Withdrawapplication(req types.WithdrawApplicationReq) (resp *types.CommonResponse, err error) {
	_, account, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.CommonResult(l.svcCtx.UserRpc.CancelJoinTenantRequest(l.ctx, &user.CancelJoinTenantReq{
		Account:     account,
		TenantCodes: req.TenantCodes,
	}))
}
