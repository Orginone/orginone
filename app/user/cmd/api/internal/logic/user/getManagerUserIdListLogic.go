package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManagerUserIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetManagerUserIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetManagerUserIdListLogic {
	return GetManagerUserIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetManagerUserIdListLogic) GetManagerUserIdList(req types.GetManagerUserIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
