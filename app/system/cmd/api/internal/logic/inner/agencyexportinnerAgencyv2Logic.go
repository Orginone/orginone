package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyexportinnerAgencyv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyexportinnerAgencyv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyexportinnerAgencyv2Logic {
	return AgencyexportinnerAgencyv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyexportinnerAgencyv2Logic) AgencyexportinnerAgencyv2(req types.AgencyExportInnerAgencyV2Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
