package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetbelowunitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetbelowunitLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetbelowunitLogic {
	return GetbelowunitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetbelowunitLogic) Getbelowunit(req types.GetBelowUnitReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindBlowUnitPageByGroupId(l.ctx, &system.FindUnitPageReq{
		GroupId: req.GroupId,
		Page: &system.PageRequest{
			Offset: (req.Current - 1) * req.Size,
			Limit:  req.Size,
			Filter: req.Name,
		},
	})
	return types.PageResult(rpcres, err)
}
