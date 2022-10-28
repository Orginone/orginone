package async

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/excel"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportMasterUnitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportMasterUnitLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportMasterUnitLogic {
	return ImportMasterUnitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportMasterUnitLogic) ImportMasterUnit(req types.ImportMasterUnitReq) (resp *types.CommonResponse, err error) {
	units, err := excel.ReadMasterUnitFromFile(req.File)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	datas := make([]model.UpdatePersonMasterData, 0)
	for _, v := range units {
		datas = append(datas, model.UpdatePersonMasterData{
			TenantCode:  tenantCode,
			PersonName:  v.LinkMan,
			PhoneNumber: v.LinkPhone,
		})
	}
	lenth := len(datas)
	for index := 0; index < lenth; index += 100 {
		endIndex := index + 100
		if lenth < endIndex {
			endIndex = lenth
		}
		buffer, err := json.Marshal(datas[index:endIndex])
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		}
		l.svcCtx.SystemRpc.UpdateMsterUnit(context.Background(), &system.CommonRpcReq{Json: string(buffer)})
	}
	//TODO 该功能原先程序未实现
	// buffer, err := json.Marshal(&schema.AsInputData{
	// 	FileName:   req.FileName,
	// 	TableName:  "人员用户表",
	// 	TCount:     int64(lenth),
	// 	FCount:     0,
	// 	Status:     1,
	// 	Type:       1,
	// 	TenantCode: req.TenantCode,
	// 	Context:    "导入成功!",
	// 	TotalTime:  0,
	// })
	// if err != nil {
	// 	return types.Failed(http.StatusInternalServerError, err)
	// }
	// l.svcCtx.SystemRpc.CreateInputRecord(context.Background(), &system.CommonRpcReq{
	// 	Json: string(buffer),
	// })
	return types.Successful(tools.GenTenantCode())
}
