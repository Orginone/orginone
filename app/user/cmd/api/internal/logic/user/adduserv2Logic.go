package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Adduserv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdduserv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Adduserv2Logic {
	return Adduserv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Adduserv2Logic) Adduserv2(req types.AddUserV2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
