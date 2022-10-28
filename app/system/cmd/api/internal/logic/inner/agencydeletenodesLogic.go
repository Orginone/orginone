package inner

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencydeletenodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencydeletenodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencydeletenodesLogic {
	return AgencydeletenodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencydeletenodesLogic) Agencydeletenodes(req types.AgencyDeleteNodesReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.AgencyRemove(l.ctx, &system.AgencyRemoveReq{Ids: tools.ArrayStrToInt64(strings.Split(req.Ids, ","))}))
}
