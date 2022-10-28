package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRoleIdListLogic {
	return GetRoleIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleIdListLogic) GetRoleIdList(req types.GetRoleIdListReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindRoleIdListByUserId(l.ctx, &system.FindRoleListByUserIdReq{
		Page: &system.PageRequest{
			Offset:      (req.Current - 1) * req.Size,
			Limit:       req.Size,
			SearchCount: true,
		},
		UserId: req.UserId,
	}))
}
