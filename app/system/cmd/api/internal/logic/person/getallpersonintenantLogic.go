package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallpersonintenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallpersonintenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallpersonintenantLogic {
	return GetallpersonintenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallpersonintenantLogic) Getallpersonintenant(req types.GetAllPersonInTenantReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindTenantPersons(l.ctx, &system.SearchPersonByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: req.Size * (req.Current - 1),
			Limit:  req.Size,
			Filter: req.FuzzyVal,
		},
		TenantCode: req.TenantCode,
		IsActivate: req.IsActivate,
		DeptId:     0,
		JobId:      0,
	}))
}
