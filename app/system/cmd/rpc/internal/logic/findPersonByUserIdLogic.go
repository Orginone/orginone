package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPersonByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPersonByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPersonByUserIdLogic {
	return &FindPersonByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *FindPersonByUserIdLogic) FindPersonByUserId(in *system.PersonByUserReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.QueryPersonInfo(l.ctx, in.UserId))
}
