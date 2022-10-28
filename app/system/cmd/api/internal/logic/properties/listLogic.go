package properties

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListReq1) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindPropertiesList(l.ctx, &system.FindPropertiesListReq{
		Page: &system.PageRequest{
			Filter: req.LikeName,
			Limit:  req.Size,
			Offset: (req.Current - 1) * req.Size,
		},
		GroupId: req.GroupId,
	}))
}
