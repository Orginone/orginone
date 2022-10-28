package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RealGroupMergeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRealGroupMergeLogic(ctx context.Context, svcCtx *svc.ServiceContext) RealGroupMergeLogic {
	return RealGroupMergeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RealGroupMergeLogic) RealGroupMerge(req types.RealGroupMergeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
