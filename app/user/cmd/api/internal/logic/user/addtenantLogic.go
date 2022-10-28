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

type AddtenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddtenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddtenantLogic {
	return AddtenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddtenantLogic) Addtenant(req types.AddTenantReq) (resp *types.CommonResponse, err error) {
	userId, account, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	_, err = l.svcCtx.UserRpc.JoinTeantByCode(l.ctx, &user.TenantCodeReq{
		UserId:     userId,
		Account:    account,
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(true)
}
