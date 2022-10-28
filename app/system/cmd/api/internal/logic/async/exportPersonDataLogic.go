package async

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

type ExportPersonDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportPersonDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExportPersonDataLogic {
	return ExportPersonDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportPersonDataLogic) ExportPersonData(req types.ExportPersonDataReq) (resp []byte, err error) {
	if tools.IsNilOrEmpty(req.TenantCode) {
		return excel.WritePersonsToFile(make([]*schema.AsPerson, 0))
	}
	rpcres, err := l.svcCtx.SystemRpc.FindTenantPersons(l.ctx, &system.SearchPersonByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: 0,
			Limit:  99999,
			Filter: "",
		},
		TenantCode: req.TenantCode,
		IsActivate: 3,
		DeptId:     0,
		JobId:      0,
	})
	if err != nil {
		return nil, err
	}
	var data *struct {
		Records []*schema.AsPerson
	}
	err = json.Unmarshal([]byte(rpcres.Json), &data)
	if err != nil {
		return nil, err
	}
	return excel.WritePersonsToFile(data.Records)
}
