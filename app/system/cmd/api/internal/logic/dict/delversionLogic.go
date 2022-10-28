package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelversionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelversionLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelversionLogic {
	return DelversionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelversionLogic) Delversion(req types.DicDelversionReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
