package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveCacheLogic {
	return RemoveCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCacheLogic) RemoveCache(req types.RemoveCacheReq) (resp *types.CommonResponse, err error) {
	return types.Successful("移除成功")
}
