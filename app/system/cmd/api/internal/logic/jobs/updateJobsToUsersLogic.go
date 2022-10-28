package jobs

import (
	"context"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJobsToUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateJobsToUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateJobsToUsersLogic {
	return UpdateJobsToUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateJobsToUsersLogic) UpdateJobsToUsers(req types.UpdateJobsToUsersReq) (resp *types.CommonResponse, err error) {
	return types.BoolResult(l.svcCtx.SystemRpc.UpdateJobUsers(l.ctx, &system.UpdateJobUsersReq{
		JobIds:    tools.ArrayStrToInt64(strings.Split(req.JobIds, ",")),
		PersonIds: tools.ArrayStrToInt64(strings.Split(req.UserIds, ",")),
	}))
}
