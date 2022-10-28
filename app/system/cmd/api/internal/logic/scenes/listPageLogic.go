package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListPageLogic {
	return ListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPageLogic) ListPage(req types.ListPageReq4) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindScenesListPage(l.ctx, &system.FindScenesListPageReq{
		Page: &system.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		TenantName: req.TenantName,
		TenantCode: req.TenantCode,
		IsDeleted:  req.IsDeleted,
		Id:         req.Id,
	}))
}
