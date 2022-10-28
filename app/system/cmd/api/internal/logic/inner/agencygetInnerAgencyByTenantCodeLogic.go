package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetInnerAgencyByTenantCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetInnerAgencyByTenantCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetInnerAgencyByTenantCodeLogic {
	return AgencygetInnerAgencyByTenantCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetInnerAgencyByTenantCodeLogic) AgencygetInnerAgencyByTenantCode(req types.AgencyGetInnerAgencyByTenantCodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
