package scenes

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/excel"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTenantExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportTenantExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExportTenantExcelLogic {
	return ExportTenantExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportTenantExcelLogic) ExportTenantExcel(req types.ExportTenantExcelReq) (resp []byte, err error) {
	rpcRes, err := l.svcCtx.SystemRpc.FindTenantByNames(l.ctx, &system.CommonRpcReq{Json: req.TenantNames})
	if err != nil {
		return nil, err
	}
	tenants := make([]*schema.AsTenant, 0)
	if err = json.Unmarshal([]byte(rpcRes.Json), &tenants); err != nil {
		return nil, err
	}
	return excel.WriteTenantToFile(tenants)
}
