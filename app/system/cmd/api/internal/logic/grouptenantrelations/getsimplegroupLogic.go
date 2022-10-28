package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetsimplegroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetsimplegroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetsimplegroupLogic {
	return GetsimplegroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetsimplegroupLogic) Getsimplegroup(req types.GetSimpleGroupReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
