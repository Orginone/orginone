package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2NodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2NodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2NodeLogic {
	return V2NodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2NodeLogic) V2Node(req types.V2NodeReq2) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
