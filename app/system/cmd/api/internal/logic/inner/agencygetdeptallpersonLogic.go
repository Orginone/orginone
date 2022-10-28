package inner

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetdeptallpersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetdeptallpersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetdeptallpersonLogic {
	return AgencygetdeptallpersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetdeptallpersonLogic) Agencygetdeptallperson(req types.AgencyGetDeptAllPersonReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInsufficientStorage, err)
	}
	return types.PageResult(l.svcCtx.SystemRpc.FindTenantPersons(l.ctx, &system.SearchPersonByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: req.Size * (req.Current - 1),
			Limit:  req.Size,
			Filter: req.FuzzyVal,
		},
		TenantCode: tenantCode,
		IsActivate: req.IsActivate,
		DeptId:     req.DeptId,
		JobId:      0,
		HasChild:   true,
	}))
}
