package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asjob"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindJobListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindJobListLogic {
	return &FindJobListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取岗位列表
func (l *FindJobListLogic) FindJobList(in *system.FindJobListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsJob.Query().Where(asjob.IsDeleted(0), asjob.TenantCode(in.TenantCode))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	jobEntitys, err := query.Order(schema.Asc(asjob.FieldSort)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, total, jobEntitys, err)
}
