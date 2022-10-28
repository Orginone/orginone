package token

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserretrieveuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserretrieveuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserretrieveuserLogic {
	return UserretrieveuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserretrieveuserLogic) Userretrieveuser(req types.UserRetrieveUserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
