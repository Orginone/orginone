package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Updatev2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatev2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Updatev2Logic {
	return Updatev2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Updatev2Logic) Updatev2(req types.UpdateV2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
