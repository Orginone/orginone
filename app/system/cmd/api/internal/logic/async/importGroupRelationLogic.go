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

type ImportGroupRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportGroupRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportGroupRelationLogic {
	return ImportGroupRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportGroupRelationLogic) ImportGroupRelation(req types.ImportGroupRelationReq) (resp *types.CommonResponse, err error) {
	groups, err := excel.ReadGroupRelationFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	userId, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	for _, v := range groups {
		v.TenantCode = tenantCode
		v.SourceGroupId = req.SourceGroupId
		v.UserId = userId
	}
	lenth := len(groups)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		buffer, err := json.Marshal(groups[index:endIndex])
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		}
		l.svcCtx.SystemRpc.AddGroupRelations(context.Background(), &system.CommonRpcReq{Json: string(buffer)})
	}
	buffer, err := json.Marshal(&schema.AsInputData{
		FileName:   req.FileName,
		TableName:  "集团间关系表",
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
