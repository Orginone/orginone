package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencydeleteUsersToAgencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencydeleteUsersToAgencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencydeleteUsersToAgencyLogic {
	return AgencydeleteUsersToAgencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencydeleteUsersToAgencyLogic) AgencydeleteUsersToAgency(req types.AgencyDeleteUsersToAgencyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
