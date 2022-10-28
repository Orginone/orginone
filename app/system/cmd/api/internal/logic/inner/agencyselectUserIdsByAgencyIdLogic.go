package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyselectUserIdsByAgencyIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyselectUserIdsByAgencyIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyselectUserIdsByAgencyIdLogic {
	return AgencyselectUserIdsByAgencyIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyselectUserIdsByAgencyIdLogic) AgencyselectUserIdsByAgencyId(req types.AgencySelectUserIdsByAgencyIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
