package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencydeletejobdeptusersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencydeletejobdeptusersLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencydeletejobdeptusersLogic {
	return AgencydeletejobdeptusersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencydeletejobdeptusersLogic) Agencydeletejobdeptusers(req types.AgencyDeleteJobDeptUsersReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.RemoveDeptJobUser(l.ctx, &system.RemoveDeptUserJobReq{
		DeptId:    req.DeptId,
		JobId:     req.JobId,
		PersonIds: []int64{req.UserIds},
	}))
}
