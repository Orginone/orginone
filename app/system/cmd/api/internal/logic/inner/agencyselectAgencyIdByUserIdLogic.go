package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyselectAgencyIdByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyselectAgencyIdByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyselectAgencyIdByUserIdLogic {
	return AgencyselectAgencyIdByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyselectAgencyIdByUserIdLogic) AgencyselectAgencyIdByUserId(req types.AgencySelectAgencyIdByUserIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
