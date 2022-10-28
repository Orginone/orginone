package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreatePasswordLogic {
	return CreatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePasswordLogic) CreatePassword(req types.CreatePasswordReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
