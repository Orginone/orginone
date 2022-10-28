package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListapplytenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListapplytenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListapplytenantLogic {
	return ListapplytenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListapplytenantLogic) Listapplytenant(req types.ListApplyTenantReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.ListApplyTenant(l.ctx, &system.ListApplyTenantReq{
		GroupId: req.GroupId,
		Flag:    req.Flag,
		Page: &system.PageRequest{
			Limit:  req.Size,
			Offset: (req.Current - 1) * req.Size,
		},
	}))
}
