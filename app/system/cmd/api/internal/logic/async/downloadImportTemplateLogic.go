package async

import (
	"context"
	"errors"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common/excel"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadImportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadImportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DownloadImportTemplateLogic {
	return DownloadImportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadImportTemplateLogic) DownloadImportTemplate(req types.DownloadImportTemplateReq) (resp []byte, err error) {
	switch req.TempType {
	//内部机构导出
	case 2:
		return excel.WriteAgencyToFile(make([]*schema.AsInnerAgency, 0))
	//岗位导出
	case 3:
		return excel.WriteJobToFile(make([]*schema.AsJob, 0))
	//租户导出
	case 4:
		return excel.WriteTenantToFile(make([]*schema.AsTenant, 0))
	//集团关系导出
	case 5:
		return excel.WriteGroupRelationToFile(make([]*model.GroupRelationData, 0))
	//单位导出
	case 6:
		return excel.WriteUnitToFile(make([]*schema.AsUnit, 0))
	// //无管理员租户导入模板
	// case 8:
	// 	return excel.WriteUnitToFile(make([]*schema.AsUnit, 0))
	// //补充管理员导入模板
	// case 9:
	// 	return excel.WriteUnitToFile(make([]*schema.AsUnit, 0))
	//主单位导出
	case 11:
		return excel.WriteMasterUnitToFile(make([]*schema.AsUnit, 0))
	}
	return make([]byte, 0), errors.New("type is not defined")
}
