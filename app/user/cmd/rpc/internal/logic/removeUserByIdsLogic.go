package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserByIdsLogic {
	return &RemoveUserByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *RemoveUserByIdsLogic) RemoveUserByIds(in *user.RemoveUserByIdsReq) (*user.CommonRpcRes, error) {
	return user.JsonResult(l.svcCtx.UserStore.AsUser.Update().Where(asuser.IsDeleted(0), asuser.IDIn(in.Ids...)).SetIsDeleted(1).Save(l.ctx))
}
