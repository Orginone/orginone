package async

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/excel"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportDeptLogic {
	return ImportDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportDeptLogic) ImportDept(req types.ImportDeptReq) (resp *types.CommonResponse, err error) {
	depts, err := excel.ReadAgencyFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	for _, v := range depts {
		v.TenantCode = tenantCode
	}
	lenth := len(depts)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		buffer, err := json.Marshal(depts[index:endIndex])
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		}
		l.svcCtx.SystemRpc.AddAgencys(context.Background(), &system.CommonRpcReq{Json: string(buffer)})
	}
	buffer, err := json.Marshal(&schema.AsInputData{
		FileName:   req.FileName,
		TableName:  "租户部门表",
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
