package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencygetnodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencygetnodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencygetnodeLogic {
	return AgencygetnodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencygetnodeLogic) Agencygetnode(req types.AgencyGetNodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
