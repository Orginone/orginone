package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApprovelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApprovelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApprovelistLogic {
	return ApprovelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApprovelistLogic) Approvelist(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
