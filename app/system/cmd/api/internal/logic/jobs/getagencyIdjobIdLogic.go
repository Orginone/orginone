package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetagencyIdjobIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetagencyIdjobIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetagencyIdjobIdLogic {
	return GetagencyIdjobIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetagencyIdjobIdLogic) GetagencyIdjobId(req types.GetAgencyIdJobIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
