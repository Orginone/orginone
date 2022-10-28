package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovetenantsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovetenantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemovetenantsLogic {
	return RemovetenantsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovetenantsLogic) Removetenants(req types.RemoveTenantsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
