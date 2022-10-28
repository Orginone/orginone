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

type ImportPersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportPersonLogic {
	return ImportPersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportPersonLogic) ImportPerson(req types.ImportPersonReq) (resp *types.CommonResponse, err error) {
	persons, err := excel.ReadPersonsFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	for _, person := range persons {
		person.TenantCode = req.TenantCode
	}
	lenth := len(persons)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		buffer, err := json.Marshal(persons[index:endIndex])
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		}
		l.svcCtx.SystemRpc.AddPersons(context.Background(), &system.CommonRpcReq{Json: string(buffer)})
	}
	buffer, err := json.Marshal(&schema.AsInputData{
		FileName:   req.FileName,
		TableName:  "人员用户表",
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
