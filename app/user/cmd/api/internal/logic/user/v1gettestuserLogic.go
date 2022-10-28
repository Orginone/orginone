package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V1gettestuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV1gettestuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) V1gettestuserLogic {
	return V1gettestuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V1gettestuserLogic) V1gettestuser(req types.V1GetTestUser) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
