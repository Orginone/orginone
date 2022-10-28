package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserbyphoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserbyphoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetuserbyphoneLogic {
	return GetuserbyphoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserbyphoneLogic) Getuserbyphone(req types.GetUserByPhoneReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
