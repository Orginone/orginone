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

type ImportTenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportTenantLogic {
	return ImportTenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportTenantLogic) ImportTenant(req types.ImportTenantReq) (resp *types.CommonResponse, err error) {
	tenantDatas, err := excel.ReadTenantFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	lenth := len(tenantDatas)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		json, err := json.Marshal(tenantDatas[index:endIndex])
		if err != nil {
			return types.Failed(http.StatusBadRequest, err)
		}
		l.svcCtx.SystemRpc.AddTenants(context.Background(), &system.CommonRpcReq{
			Json: string(json),
		})
	}
	buffer, err := json.Marshal(&schema.AsInputData{
		FileName:   req.FileName,
		TableName:  "集团单位关系表",
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
