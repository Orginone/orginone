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

type AgencyaddinneruserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyaddinneruserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyaddinneruserLogic {
	return AgencyaddinneruserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyaddinneruserLogic) Agencyaddinneruser(req types.AgencyAddInnerUserReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.AddInnerUser(l.ctx, &system.AddInnerUserReq{
		ParentId:   req.ParentId,
		PersonIds:  tools.ArrayStrToInt64(strings.Split(req.UserIds, ",")),
		TenantCode: req.TenantCode,
		AgencyName: req.AgencyName,
		AgencyCode: req.AgencyCode,
	}))
}
