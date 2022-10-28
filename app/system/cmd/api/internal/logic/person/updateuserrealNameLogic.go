package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuserrealNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuserrealNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateuserrealNameLogic {
	return UpdateuserrealNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuserrealNameLogic) UpdateuserrealName(req types.UpdateUserRealNameReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.UpdateUserName(l.ctx, &system.UpdateUserName{
		UserId:   req.UserId,
		IdCard:   req.IdCard,
		UserName: req.RealName,
		UserCode: req.UserCode,
	}))
}
