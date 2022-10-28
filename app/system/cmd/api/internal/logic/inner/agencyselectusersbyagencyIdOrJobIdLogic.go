package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyselectusersbyagencyIdOrJobIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyselectusersbyagencyIdOrJobIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyselectusersbyagencyIdOrJobIdLogic {
	return AgencyselectusersbyagencyIdOrJobIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyselectusersbyagencyIdOrJobIdLogic) AgencyselectusersbyagencyIdOrJobId(req types.AgencySelectUsersByAgencyIdOrJobIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
