package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteacttodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteacttodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteacttodoLogic {
	return DeleteacttodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteacttodoLogic) Deleteacttodo(req types.DeleteacttodoReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
