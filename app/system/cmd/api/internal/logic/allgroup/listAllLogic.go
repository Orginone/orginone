package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListAllLogic {
	return ListAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllLogic) ListAll(req types.ListAllReq1) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindAllGroupListAll(l.ctx, &system.FindAllGroupListAllReq{
		Page: &system.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		Type:       req.Type,
		GroupCode:  req.GroupCode,
		GroupName:  req.GroupName,
		TenantCode: req.TenantCode,
		IsDeleted:  req.IsDeleted,
		Id:         req.Id,
	}))
}
