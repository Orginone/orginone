package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitAllLogic {
	return SubmitAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitAllLogic) SubmitAll(req types.SubmitAllReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
