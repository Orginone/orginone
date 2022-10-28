package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asjob"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJobUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddJobUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJobUserLogic {
	return &AddJobUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增岗位并分配人员
func (l *AddJobUserLogic) AddJobUser(in *system.AddJobUserReq) (*system.CommonRpcRes, error) {
	isExist, err := l.svcCtx.SystemStore.AsJob.Query().Where(asjob.TenantCode(in.TenantCode), asjob.JobName(in.JobName), asjob.IsDeleted(0)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if isExist {
		return &system.CommonRpcRes{}, errors.New("job is exist by jobname:" + in.JobName)
	}
	return system.JsonResult(l.svcCtx.SystemStore.AsJob.Create().SetJobName(in.JobName).SetTenantCode(in.TenantCode).SetIsDeleted(0).AddPersonIDs(in.PersonIds...).Save(l.ctx))
}
