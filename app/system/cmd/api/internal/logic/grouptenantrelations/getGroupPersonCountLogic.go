package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupPersonCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupPersonCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetGroupPersonCountLogic {
	return GetGroupPersonCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupPersonCountLogic) GetGroupPersonCount(req types.GetGroupPersonCountReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
