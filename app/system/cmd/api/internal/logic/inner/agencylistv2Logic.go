package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type Agencylistv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencylistv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Agencylistv2Logic {
	return Agencylistv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Agencylistv2Logic) Agencylistv2(req types.AgencyListV2Req) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindInnerAgencyList(l.ctx, &system.FindInnerAgencyListReq{
		Page: &system.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		AgencyName: req.FuzzyVal,
		TenantCode: req.TenantCode,
	}))
}
