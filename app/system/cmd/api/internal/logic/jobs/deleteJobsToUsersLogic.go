package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJobsToUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteJobsToUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteJobsToUsersLogic {
	return DeleteJobsToUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteJobsToUsersLogic) DeleteJobsToUsers(req types.DeleteJobsToUsersReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
