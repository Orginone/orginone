package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetjobsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetjobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetjobsLogic {
	return AgencygetjobsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetjobsLogic) Agencygetjobs(req types.AgencyGetJobsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
