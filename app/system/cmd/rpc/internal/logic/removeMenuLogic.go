package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asmenu"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMenuLogic {
	return &RemoveMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除菜单
func (l *RemoveMenuLogic) RemoveMenu(in *system.RemoveMenuReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsMenu.Update().Where(asmenu.IDIn(in.MenuIds...)).SetIsDeleted(1).Save(l.ctx))
}
