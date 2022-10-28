package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Updatev1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatev1Logic(ctx context.Context, svcCtx *svc.ServiceContext) Updatev1Logic {
	return Updatev1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Updatev1Logic) Updatev1(req types.UpdateReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
