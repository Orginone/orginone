package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyadddeptbyappidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyadddeptbyappidLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyadddeptbyappidLogic {
	return AgencyadddeptbyappidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyadddeptbyappidLogic) Agencyadddeptbyappid(req types.AgencyAddDeptByAppidReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
