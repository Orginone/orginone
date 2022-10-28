package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2detailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2detailLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2detailLogic {
	return V2detailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2detailLogic) V2detail(req types.V2DetailReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
