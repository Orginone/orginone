package menus

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ButtonsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewButtonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ButtonsLogic {
	return ButtonsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ButtonsLogic) Buttons(req types.ButtonsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
