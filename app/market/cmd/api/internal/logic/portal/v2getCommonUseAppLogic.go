package portal

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2getCommonUseAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2getCommonUseAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2getCommonUseAppLogic {
	return V2getCommonUseAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2getCommonUseAppLogic) V2getCommonUseApp(req types.V2GetCommonUseAppReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
