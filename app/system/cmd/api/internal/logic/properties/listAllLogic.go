package properties

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListAllLogic {
	return ListAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllLogic) ListAll(req types.ListAllReq2) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
