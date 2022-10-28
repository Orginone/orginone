package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveuserLogic {
	return RemoveuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveuserLogic) Removeuser(req types.RemoveUserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
