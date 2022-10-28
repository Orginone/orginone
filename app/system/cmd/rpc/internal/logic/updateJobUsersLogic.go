package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"
	"orginone/common/schema/asworkingdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJobUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateJobUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJobUsersLogic {
	return &UpdateJobUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量修改岗位并分配人员
func (l *UpdateJobUsersLogic) UpdateJobUsers(in *system.UpdateJobUsersReq) (*system.CommonRpcRes, error) {
	return system.JsonResult(l.svcCtx.SystemStore.AsPerson.Update().
		Where(
			asperson.IDIn(in.PersonIds...),
			asperson.IsDeleted(0),
			asperson.Not(
				asperson.HasUserxWith(
					asuser.HasWorkingDatasWith(
						asworkingdata.Type(3),
						asworkingdata.IsDeleted(0),
					),
				),
			),
		).ClearJobs().AddJobIDs(in.JobIds...).Save(l.ctx))
}
