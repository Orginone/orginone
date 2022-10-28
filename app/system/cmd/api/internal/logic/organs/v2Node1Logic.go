package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2nodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2Node1Logic(ctx context.Context, svcCtx *svc.ServiceContext) V2nodeLogic {
	return V2nodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2nodeLogic) V2Node1(req types.V2NodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
