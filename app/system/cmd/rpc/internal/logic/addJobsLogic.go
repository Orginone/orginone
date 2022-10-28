package logic

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asjob"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJobsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddJobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJobsLogic {
	return &AddJobsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 岗位导入
func (l *AddJobsLogic) AddJobs(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	l.ctx = context.Background()
	jobs := make([]*schema.AsJob, 0)
	err := json.Unmarshal([]byte(in.Json), &jobs)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	jobNames := make([]string, 0)
	for _, v := range jobs {
		jobNames = append(jobNames, v.JobName)
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	jobEntityNames, err := tx.AsJob.Query().Where(asjob.IsDeleted(0), asjob.TenantCode(jobs[0].TenantCode), asjob.JobNameIn(jobNames...)).Select(asjob.FieldJobName).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	bluks := make([]*schema.AsJobCreate, 0)
	for _, v := range jobs {
		if tools.ArrIndexOf(jobEntityNames, v.JobName) < 0 {
			bluks = append(bluks, l.svcCtx.SystemStore.AsJob.Create().SetJobName(v.JobName).SetTenantCode(v.TenantCode))
		}
	}
	_, err = tx.AsJob.CreateBulk(bluks...).Save(l.ctx)
	return system.Result("true", tx.Commit())
}
