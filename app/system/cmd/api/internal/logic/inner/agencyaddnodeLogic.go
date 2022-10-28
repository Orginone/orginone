package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyaddnodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyaddnodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyaddnodeLogic {
	return AgencyaddnodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyaddnodeLogic) Agencyaddnode(req types.AgencyAddNodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
