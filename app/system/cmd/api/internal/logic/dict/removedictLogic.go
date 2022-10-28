package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovedictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovedictLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemovedictLogic {
	return RemovedictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovedictLogic) Removedict(req types.DicRemovedictReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
