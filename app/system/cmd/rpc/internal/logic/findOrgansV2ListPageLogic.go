package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asunit"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrgansV2ListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrgansV2ListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrgansV2ListPageLogic {
	return &FindOrgansV2ListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询单位
func (l *FindOrgansV2ListPageLogic) FindOrgansV2ListPage(in *system.FindOrgansV2ListPageReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.UnitNameContains(in.UnitName),
		asunit.SocialCreditCodeContains(in.SocialCreditCode))
	if in.Id != -1 {
		query.Where(asunit.ID(in.Id))
	}
	if in.IsDeleted != -1 {
		query.Where(asunit.IsDeleted(in.IsDeleted))
	}
	if in.TenantId != -1 {
		query.Where(asunit.TenantID(in.TenantId))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	unitEntitys, err := query.Order(schema.Asc(asunit.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(count), unitEntitys, err)
}
