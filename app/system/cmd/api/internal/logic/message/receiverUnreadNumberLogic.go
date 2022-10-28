package message

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiverUnreadNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiverUnreadNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReceiverUnreadNumberLogic {
	return ReceiverUnreadNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiverUnreadNumberLogic) ReceiverUnreadNumber(req types.Nil) (resp *types.CommonResponse, err error) {
	return types.JsonResult(false, nil)
}
