package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/baseinfoadministrativeareaall"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAdministrativeTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAdministrativeTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAdministrativeTreeLogic {
	return &FindAdministrativeTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取行政区域树
func (l *FindAdministrativeTreeLogic) FindAdministrativeTree(in *system.Nil) (*system.CommonRpcRes, error) {
	administrativeEntitys, err := l.svcCtx.SystemStore.Baseinfoadministrativeareaall.Query().Where(baseinfoadministrativeareaall.IsDeleted(0)).All(l.ctx)
	return system.JsonResult(loopMakeAdminTreeByIds(l.ctx, administrativeEntitys, 0), err)
}

//make tree
func loopMakeAdminTreeByIds(ctx context.Context, array []*schema.Baseinfoadministrativeareaall, parentID int64) []*model.AdministrativeTree {
	adminTrees := make([]*model.AdministrativeTree, 0)
	for _, item := range array {
		if item.Pid == parentID {
			adminTree := new(model.AdministrativeTree)
			adminTree.Lable = item.Name
			adminTree.Value = item.Name
			adminTree.Children = loopMakeAdminTreeByIds(ctx, array, item.ID)
			adminTrees = append(adminTrees, adminTree)
		}
	}
	return adminTrees
}
