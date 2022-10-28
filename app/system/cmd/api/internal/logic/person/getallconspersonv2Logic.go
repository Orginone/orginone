package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Getallconspersonv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallconspersonv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Getallconspersonv2Logic {
	return Getallconspersonv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Getallconspersonv2Logic) Getallconspersonv2(req types.GetAllConsPersonV2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
