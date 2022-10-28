package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveProgressLogic {
	return RemoveProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveProgressLogic) RemoveProgress(req types.RemoveProgressReq) (resp *types.CommonResponse, err error) {
	return types.Successful("移除成功")
}
