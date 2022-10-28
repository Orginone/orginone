package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyselectUsersByAgencyIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyselectUsersByAgencyIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyselectUsersByAgencyIdLogic {
	return AgencyselectUsersByAgencyIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyselectUsersByAgencyIdLogic) AgencyselectUsersByAgencyId(req types.AgencySelectUsersByAgencyIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
