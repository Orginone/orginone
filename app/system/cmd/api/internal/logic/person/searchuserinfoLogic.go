package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchuserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchuserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchuserinfoLogic {
	return SearchuserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchuserinfoLogic) Searchuserinfo(req types.SearchUserInfoReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindTenantPersons(l.ctx, &system.SearchPersonByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: req.Size * (req.Current - 1),
			Limit:  req.Size,
			Filter: req.FuzzyVal,
		},
		TenantCode: req.TenantCode,
		IsActivate: req.IsActivate,
		DeptId:     req.DeptId,
		JobId:      req.JobId,
	}))
}
