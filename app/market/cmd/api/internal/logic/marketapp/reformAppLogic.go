package marketapp

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReformAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReformAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReformAppLogic {
	return ReformAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReformAppLogic) ReformApp(req types.ReformAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
