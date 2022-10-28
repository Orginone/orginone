package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllGroupByCodeAndTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllGroupByCodeAndTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllGroupByCodeAndTypeLogic {
	return &FindAllGroupByCodeAndTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据组织编号和类型
func (l *FindAllGroupByCodeAndTypeLogic) FindAllGroupByCodeAndType(in *system.SearchAllGroupReq) (*system.CommonRpcRes, error) {
	groupEntitys, err := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.TenantCode(in.TenantCode),
		asallgroup.IsDeleted(0), asallgroup.Type(in.GroupType)).All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(groupEntitys, err)
}
