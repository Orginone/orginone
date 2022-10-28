package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawGroupAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawGroupAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) WithdrawGroupAuthorityLogic {
	return WithdrawGroupAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WithdrawGroupAuthorityLogic) WithdrawGroupAuthority(req types.WithdrawGroupAuthorityReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
