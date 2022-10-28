package async

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/excel"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportJobLogic {
	return ImportJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportJobLogic) ImportJob(req types.ImportJobReq) (resp *types.CommonResponse, err error) {
	jobs, err := excel.ReadJobFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	for _, v := range jobs {
		v.TenantCode = req.TenantCode
	}
	lenth := len(jobs)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		buffer, err := json.Marshal(jobs[index:endIndex])
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		}
		l.svcCtx.SystemRpc.AddJobs(context.Background(), &system.CommonRpcReq{Json: string(buffer)})
	}
	buffer, err := json.Marshal(&schema.AsInputData{
		FileName:   req.FileName,
		TableName:  "租户岗位表",
		TCount:     int64(lenth),
		FCount:     0,
		Status:     1,
		Type:       1,
		TenantCode: req.TenantCode,
		Context:    "导入成功!",
		TotalTime:  0,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	l.svcCtx.SystemRpc.CreateInputRecord(context.Background(), &system.CommonRpcReq{
		Json: string(buffer),
	})
	return types.Successful(tools.GenTenantCode())
}
