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

type ImportTenantRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportTenantRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportTenantRelationLogic {
	return ImportTenantRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportTenantRelationLogic) ImportTenantRelation(req types.ImportTenantRelationReq) (resp *types.CommonResponse, err error) {
	tenantRelationDatas, err := excel.ReadTenantRelationFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	lenth := len(tenantRelationDatas)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		l.svcCtx.SystemRpc.UpdateTenantRelations(context.Background(), &system.UpdateTenantRelationsReq{
			TenantRelationDatas: tenantRelationDatas[index:endIndex],
			SourceGroupId:       req.SourceGroupId,
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
