package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByUserIdLogic {
	return &FindUserByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *FindUserByUserIdLogic) FindUserByUserId(in *system.PrimaryKeyReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsUser.Get(l.ctx, in.Id))
}
