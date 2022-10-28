package async

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/common/excel"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportGroupUnitDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportGroupUnitDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExportGroupUnitDataLogic {
	return ExportGroupUnitDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportGroupUnitDataLogic) ExportGroupUnitData(req []int64) (resp []byte, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindUnitByIds(l.ctx, &system.FindUnitByIdsReq{
		UnitIdList: req,
	})
	if err != nil {
		return nil, err
	}
	Records := make([]*schema.AsUnit, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &Records)
	if err != nil {
		return nil, err
	}
	return excel.WriteUnitToFile(Records)
}
