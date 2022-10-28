package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportTenantIconLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImportTenantIconLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportTenantIconLogic {
	return &ImportTenantIconLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 租户图片批量导入
func (l *ImportTenantIconLogic) ImportTenantIcon(in *system.ImportTenantIconReq) (*system.CommonRpcRes, error) {
	icons := make([]string, 0)
	bulks := make([]*schema.AsTenantIconCreate, 0)
	for _, v := range in.Urls {
		isExist := false
		for _, icon := range icons {
			if icon == v {
				isExist = true
				break
			}
		}
		if !isExist {
			bulks = append(bulks, l.svcCtx.SystemStore.AsTenantIcon.Create().SetTenantCode(in.TenantCode).SetIcon(v))
			icons = append(icons, v)
		}
	}
	return system.JsonResult(l.svcCtx.SystemStore.AsTenantIcon.CreateBulk(bulks...).Save(l.ctx))
}
