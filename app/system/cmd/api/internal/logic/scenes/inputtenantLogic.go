package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InputtenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInputtenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) InputtenantLogic {
	return InputtenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InputtenantLogic) Inputtenant(req types.InputTenantReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
