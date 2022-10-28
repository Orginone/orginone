package inner

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/excel"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyexportInnerAgencyDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyexportInnerAgencyDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyexportInnerAgencyDataLogic {
	return AgencyexportInnerAgencyDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyexportInnerAgencyDataLogic) AgencyexportInnerAgencyData(req types.AgencyExportInnerAgencyDataReq) (resp []byte, err error) {
	rpcRes, err := l.svcCtx.SystemRpc.FindInnerAgencyList(l.ctx, &system.FindInnerAgencyListReq{
		Ids:        tools.ArrayStrToInt64(req.InnerAgencyList),
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return nil, err
	}
	agencys := make([]*schema.AsInnerAgency, 0)
	if err = json.Unmarshal([]byte(rpcRes.Json), &agencys); err != nil {
		return nil, err
	}
	return excel.WriteAgencyToFile(agencys)
}
