package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsDistributedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsDistributedLogic(ctx context.Context, svcCtx *svc.ServiceContext) IsDistributedLogic {
	return IsDistributedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsDistributedLogic) IsDistributed(req types.IsDistributedReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
