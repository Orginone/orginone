package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJobsToUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddJobsToUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddJobsToUsersLogic {
	return AddJobsToUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddJobsToUsersLogic) AddJobsToUsers(req types.AddJobsToUsersReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
