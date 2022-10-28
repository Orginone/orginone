package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyimportinnerAgencyv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyimportinnerAgencyv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyimportinnerAgencyv2Logic {
	return AgencyimportinnerAgencyv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyimportinnerAgencyv2Logic) AgencyimportinnerAgencyv2(req types.AgencyImportInnerAgencyV2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
