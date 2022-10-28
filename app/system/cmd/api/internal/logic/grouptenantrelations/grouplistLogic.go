package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrouplistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrouplistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrouplistLogic {
	return GrouplistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrouplistLogic) Grouplist(req types.GroupListReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindGroupListByRelations(l.ctx, &system.FindGroupListByRelationsReq{
		Page: &system.PageRequest{
			Offset: (req.Current - 1) * req.Size,
			Limit:  req.Size,
		},
		GroupId: req.GroupId,
	}))
}
