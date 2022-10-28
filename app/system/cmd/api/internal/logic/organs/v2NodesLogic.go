package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2NodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2NodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2NodesLogic {
	return V2NodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2NodesLogic) V2Nodes(req types.V2NodesReq2) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
