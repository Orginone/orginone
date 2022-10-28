package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetjobspageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetjobspageLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetjobspageLogic {
	return AgencygetjobspageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetjobspageLogic) Agencygetjobspage(req types.AgencyGetJobsPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
