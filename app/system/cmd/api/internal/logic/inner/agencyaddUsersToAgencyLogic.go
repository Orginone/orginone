package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyaddUsersToAgencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyaddUsersToAgencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyaddUsersToAgencyLogic {
	return AgencyaddUsersToAgencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyaddUsersToAgencyLogic) AgencyaddUsersToAgency(req types.AgencyAddUsersToAgencyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
