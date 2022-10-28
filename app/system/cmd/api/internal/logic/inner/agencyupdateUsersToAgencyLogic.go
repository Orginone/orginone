package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyupdateUsersToAgencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyupdateUsersToAgencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyupdateUsersToAgencyLogic {
	return AgencyupdateUsersToAgencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyupdateUsersToAgencyLogic) AgencyupdateUsersToAgency(req types.AgencyUpdateUsersToAgencyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
