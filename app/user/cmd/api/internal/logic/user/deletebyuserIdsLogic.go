package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletebyuserIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletebyuserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeletebyuserIdsLogic {
	return DeletebyuserIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletebyuserIdsLogic) DeletebyuserIds(req types.DeleteByUserIdsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
