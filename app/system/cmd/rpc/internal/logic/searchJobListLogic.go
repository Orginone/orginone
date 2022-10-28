package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asjob"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchJobListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchJobListLogic {
	return &SearchJobListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 模糊搜索岗位列表
func (l *SearchJobListLogic) SearchJobList(in *system.SearchJobListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsJob.Query().Where(asjob.IsDeleted(0), asjob.TenantCode(in.TenantCode), asjob.JobNameContains(in.Page.Filter))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, nil
	}
	jobEntitys, err := query.Order(schema.Asc(asjob.FieldSort)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, total, jobEntitys, err)
}
