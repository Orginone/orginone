package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveInfoLogic {
	return RemoveInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveInfoLogic) RemoveInfo(req types.RemoveInfoReq) (resp *types.CommonResponse, err error) {
	return types.Successful("移除成功")
}
