package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type Listv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Listv2Logic {
	return Listv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Listv2Logic) Listv2(req types.ListV2Req) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindJobList(l.ctx, &system.FindJobListReq{
		Page: &system.PageRequest{
			Offset: (req.Current - 1) * req.Size,
			Limit:  req.Size,
		},
		TenantCode: req.TenantCode,
	})
	return types.PageResult(rpcres, err)
}
