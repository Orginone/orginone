package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawTenantAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawTenantAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) WithdrawTenantAuthorityLogic {
	return WithdrawTenantAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WithdrawTenantAuthorityLogic) WithdrawTenantAuthority(req types.WithdrawTenantAuthorityReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
