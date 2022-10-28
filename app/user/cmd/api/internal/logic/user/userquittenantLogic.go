package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserquittenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserquittenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserquittenantLogic {
	return UserquittenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserquittenantLogic) Userquittenant(req types.UserQuitTenantReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
