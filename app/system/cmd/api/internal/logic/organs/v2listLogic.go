package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2listLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2listLogic {
	return V2listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2listLogic) V2list(req types.V2ListReq) (resp *types.CommonResponse, err error) {
	return
}
