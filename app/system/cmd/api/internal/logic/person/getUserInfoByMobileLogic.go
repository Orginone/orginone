package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoByMobileLogic {
	return GetUserInfoByMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByMobileLogic) GetUserInfoByMobile(req types.GetUserInfoByMobileReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
