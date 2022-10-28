package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SinglelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSinglelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) SinglelistLogic {
	return SinglelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SinglelistLogic) Singlelist(req types.SingleListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
