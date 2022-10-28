package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asjob"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJobUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateJobUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJobUserLogic {
	return &UpdateJobUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改岗位并分配人员
func (l *UpdateJobUserLogic) UpdateJobUser(in *system.UpdateJobUserReq) (*system.CommonRpcRes, error) {
	jobEntity, err := l.svcCtx.SystemStore.AsJob.Query().Where(asjob.ID(in.JobId), asjob.IsDeleted(0)).WithPersons().First(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, errors.New("job not found !" + err.Error())
	}
	update := jobEntity.Update()
	if !tools.IsNilOrEmpty(in.JobName) {
		update.SetJobName(in.JobName)
	}
	if in.Sort != -1 {
		update.SetSort(in.Sort)
	}
	return system.JsonResult(update.ClearPersons().AddPersonIDs(in.PersonIds...).Save(l.ctx))
}
