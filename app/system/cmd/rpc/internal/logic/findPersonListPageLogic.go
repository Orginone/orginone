package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asperson"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPersonListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPersonListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPersonListPageLogic {
	return &FindPersonListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询人员
func (l *FindPersonListPageLogic) FindPersonListPage(in *system.FindPersonListPageReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsPerson.Query().Where(asperson.PhoneNumberContains(in.PhoneNumber),
		asperson.RealNameContains(in.RealName), asperson.TenantCodeContains(in.TenantCode))
	if in.Id != -1 {
		query.Where(asperson.ID(in.Id))
	}
	if in.IsDeleted != -1 {
		query.Where(asperson.IsDeleted(in.IsDeleted))
	} else {
		query.Where(asperson.IsDeleted(0))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	personEntitys, err := query.Order(schema.Asc(asperson.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(count), personEntitys, err)
}
