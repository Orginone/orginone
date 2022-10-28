package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyaddnodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyaddnodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyaddnodesLogic {
	return AgencyaddnodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyaddnodesLogic) Agencyaddnodes(req types.AgencyAddNodesReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
