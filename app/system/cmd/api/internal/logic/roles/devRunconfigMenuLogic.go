package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DevRunconfigMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDevRunconfigMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) DevRunconfigMenuLogic {
	return DevRunconfigMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DevRunconfigMenuLogic) DevRunconfigMenu(req types.DevRunConfigMenuReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
