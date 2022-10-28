package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteallacttodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteallacttodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteallacttodoLogic {
	return DeleteallacttodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteallacttodoLogic) Deleteallacttodo(req types.DeleteallacttodoReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
