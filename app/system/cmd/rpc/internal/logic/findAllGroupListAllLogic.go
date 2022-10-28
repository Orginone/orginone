package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllGroupListAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllGroupListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllGroupListAllLogic {
	return &FindAllGroupListAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询集团
func (l *FindAllGroupListAllLogic) FindAllGroupListAll(in *system.FindAllGroupListAllReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.GroupNameContains(in.GroupName),
		asallgroup.TenantCodeContains(in.TenantCode), asallgroup.GroupCodeContains(in.GroupCode))
	if in.Id != -1 {
		query.Where(asallgroup.ID(in.Id))
	}
	if in.IsDeleted != -1 {
		query.Where(asallgroup.IsDeleted(in.IsDeleted))
	} else {
		query.Where(asallgroup.IsDeleted(0))
	}
	if in.Type != -1 {
		query.Where(asallgroup.Type(in.Type))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	allGroupEntitys, err := query.Order(schema.Asc(asallgroup.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(count), allGroupEntitys, err)
}
